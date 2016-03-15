package user

import (
	"github.com/huandu/facebook"
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	"server/credentials"
)

type AuthReq struct {
	Service string `json:"service"`
	Token   string `json:"token"`
}

var app = facebook.New(credentials.FACEBOOK_CLIENT_ID, credentials.FACEBOOK_CLIENT_SECRET)

func VerifyToken(c context.Context, req AuthReq) (bool, error) {
	switch req.Service {
	case "fb":
		return VerifyFBToken(c, req.Token)
	}
	return false, nil
}

func VerifyFBToken(c context.Context, token string) (bool, error) {
	session := app.Session(token)
	session.HttpClient = urlfetch.Client(c)

	_, err := session.Get("/me", nil)

	if err != nil {
		return false, err
	}

	return true, nil
}
