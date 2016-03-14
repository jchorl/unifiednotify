package server

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"server/auth"
	"server/constants"
	"server/service/gmail"
	"server/tokenstore"
)

func init() {
	http.HandleFunc("/", authGoogle)
	http.HandleFunc("/googleoauthcallback", authGoogleCallback)
	http.HandleFunc("/doneauth", doneAuth)
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
	}

	// save the token
	err = tokenstore.SaveToken(c, userID, tokenstore.GOOGLE_TOKEN, tok)
	if err != nil {
		log.Errorf(c, err.Error())
	}

	// redirect
	url := constants.BASE_URL + "/doneauth"
	http.Redirect(w, r, url, http.StatusFound)
}

func doneAuth(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	tkn, err := tokenstore.GetToken(c, "josh", tokenstore.GOOGLE_TOKEN)
	if err != nil {
		log.Errorf(c, err.Error())
	}
	notifications, err := gmail.GetNotifications(c, tkn)
	if err != nil {
		log.Errorf(c, err.Error())
	}
	for _, notification := range notifications {
		fmt.Fprint(w, notification)
	}
}
