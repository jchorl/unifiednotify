package server

import (
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"server/constants"
	"server/service/auth"
	"server/service/gmail"
	"server/tokenstore"
	"server/user"
)

func init() {
	http.HandleFunc("/auth", authUser)
	http.HandleFunc("/authgoogle", authGoogle)
	http.HandleFunc("/googleoauthcallback", authGoogleCallback)
	http.HandleFunc("/donegoogleauth", doneGoogleAuth)
}

func authUser(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	decoder := json.NewDecoder(r.Body)
	var req user.AuthReq
	err := decoder.Decode(&req)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}
	log.Infof(c, "Attempting to authenticate token %s\n", req.Token)
	authenticated, err := user.VerifyToken(c, req)
	log.Infof(c, "%t\n", authenticated)
}

func authGoogle(w http.ResponseWriter, r *http.Request) {
	url := auth.GetAuthUrl("josh")
	http.Redirect(w, r, url, http.StatusFound)
}

func authGoogleCallback(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	// TODO: handle the case where error=access_denied
	code := r.FormValue("code")
	userID := r.FormValue("state")
	tok, err := auth.GetToken(c, code)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}

	// save the token
	err = tokenstore.SaveToken(c, userID, tokenstore.GOOGLE_TOKEN, tok)
	if err != nil {
		log.Errorf(c, err.Error())
		panic(err)
	}

	// redirect
	url := constants.BASE_URL + "/donegoogleauth"
	http.Redirect(w, r, url, http.StatusFound)
}

func doneGoogleAuth(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	tkn, err := tokenstore.GetToken(c, "josh", tokenstore.GOOGLE_TOKEN)
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
