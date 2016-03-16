package usermiddleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"server/credentials"
	"strings"
)

func NewAuthHandleFunc(handler func(http.ResponseWriter, *http.Request, string)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := appengine.NewContext(r)
		log.Debugf(c, "starting auth")
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			log.Debugf(c, "contains auth header")
			split := strings.SplitN(authHeader, " ", 2)
			for idx, str := range split {
				log.Debugf(c, "%d: %s", idx, str)
			}
			if len(split) == 2 {
				log.Debugf(c, "Unparsed: %s", split[1])
				token, err := jwt.Parse(split[1], func(token *jwt.Token) (interface{}, error) {
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
			}
		}
		http.Error(w, "User is not authorized to view this page. Please log in and ensure you have correct permissions.", http.StatusUnauthorized)
	}
}
