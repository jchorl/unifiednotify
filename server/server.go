package server

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"server/constants"
	"server/credentials"
	"server/service/auth"
	"server/service/gmail"
	"server/tokenstore"
	"server/user"
	"server/user/usermiddleware"
)

func init() {
	http.HandleFunc("/auth", authUser)
	http.HandleFunc("/auth/gmail/init", usermiddleware.NewAuth(authGmailInit))
	http.HandleFunc("/auth/gmail/callback", usermiddleware.NewAuth(authGmailCallback))
	http.HandleFunc("/notifications", usermiddleware.NewAuth(notifications))
}

func authUser(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	decoder := json.NewDecoder(r.Body)
	var req user.AuthReq
	log.Debugf(c, "received request to auth user")
	err := decoder.Decode(&req)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}
	userId, err := user.GetUserId(c, req)
	log.Debugf(c, "authed with userId: %s", userId)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["userId"] = userId
	tokenStr, err := token.SignedString([]byte(credentials.JWT_SIGNING_KEY))
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}
	log.Debugf(c, "Token: %s", tokenStr)
	cookie := http.Cookie{
		Name:   "auth",
		Value:  tokenStr,
		MaxAge: 604800,
		// TODO: uncomment for prod
		// Secure:   true,
		// HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func authInit(w http.ResponseWriter, r *http.Request, userId string, service string) {
	url := auth.GetAuthUrl(userId, service)
	http.Redirect(w, r, url, http.StatusFound)
}

func authCallback(w http.ResponseWriter, r *http.Request, userId string, service string) {
	c := appengine.NewContext(r)
	// TODO: handle the case where error=access_denied
	code := r.FormValue("code")
	tok, err := auth.GetToken(c, code, service)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}

	// save the token
	err = tokenstore.SaveToken(c, userId, service, tok)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}

	// redirect
	url := constants.BASE_URL
	http.Redirect(w, r, url, http.StatusFound)
}

func authGmailInit(w http.ResponseWriter, r *http.Request, userId string) {
	authInit(w, r, userId, constants.GMAIL_SERVICE)
}

func authGmailCallback(w http.ResponseWriter, r *http.Request, userId string) {
	authCallback(w, r, userId, constants.GMAIL_SERVICE)
}

func notifications(w http.ResponseWriter, r *http.Request, userId string) {
	c := appengine.NewContext(r)
	tkn, err := tokenstore.GetToken(c, userId, constants.GMAIL_SERVICE)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}
	notifications, err := gmail.GetNotifications(c, tkn)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(notifications)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}
}
