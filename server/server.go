package server

import (
	"encoding/json"
	"fmt"
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
	http.HandleFunc("/authgoogle", usermiddleware.NewAuth(authGoogle))
	http.HandleFunc("/googleoauthcallback", usermiddleware.NewAuth(authGoogleCallback))
	http.HandleFunc("/donegoogleauth", usermiddleware.NewAuth(doneGoogleAuth))
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

func authGoogle(w http.ResponseWriter, r *http.Request, userId string) {
	url := auth.GetAuthUrl(userId)
	http.Redirect(w, r, url, http.StatusFound)
}

func authGoogleCallback(w http.ResponseWriter, r *http.Request, userId string) {
	c := appengine.NewContext(r)
	// TODO: handle the case where error=access_denied
	code := r.FormValue("code")
	tok, err := auth.GetToken(c, code)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}

	// save the token
	err = tokenstore.SaveToken(c, userId, tokenstore.GOOGLE_TOKEN, tok)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}

	// redirect
	url := constants.BASE_URL + "/donegoogleauth"
	http.Redirect(w, r, url, http.StatusFound)
}

func doneGoogleAuth(w http.ResponseWriter, r *http.Request, userId string) {
	c := appengine.NewContext(r)
	tkn, err := tokenstore.GetToken(c, userId, tokenstore.GOOGLE_TOKEN)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}
	notifications, err := gmail.GetNotifications(c, tkn)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}
	for _, notification := range notifications {
		fmt.Fprint(w, notification)
	}
}
