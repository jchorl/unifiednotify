package usermiddleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"server/credentials"
)

func NewReturnIfAuthd(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := appengine.NewContext(r)
		log.Debugf(c, "starting auth")
		cookie, err := r.Cookie("auth")
		if err == nil {
			log.Debugf(c, "Unparsed: %s", cookie.Value)
			token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
				}
				return []byte(credentials.JWT_SIGNING_KEY), nil
			})
			if err == nil && token.Valid {
				log.Debugf(c, "parsed token")
				return
			} else if err != nil {
				log.Debugf(c, "error parsing")
				log.Errorf(c, err.Error())
			}
		} else {
			log.Debugf(c, err.Error())
		}
		handler(w, r)
	}
}

func NewAuth(handler func(http.ResponseWriter, *http.Request, string)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := appengine.NewContext(r)
		log.Debugf(c, "starting auth")
		cookie, err := r.Cookie("auth")
		if err == nil {
			log.Debugf(c, "Unparsed: %s", cookie.Value)
			token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
				}
				return []byte(credentials.JWT_SIGNING_KEY), nil
			})
			if err == nil && token.Valid {
				log.Debugf(c, "parsed token")
				if userId, ok := token.Claims["userId"].(string); ok {
					log.Debugf(c, "User ID: %s", userId)
					handler(w, r, userId)
					return
				}
			} else if err != nil {
				log.Debugf(c, "error parsing")
				log.Errorf(c, err.Error())
			}
		} else {
			log.Debugf(c, err.Error())
		}
		http.Error(w, "User is not authorized to view this page. Please log in and ensure you have correct permissions.", http.StatusUnauthorized)
	}
}
