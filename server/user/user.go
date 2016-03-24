package user

import (
	"errors"
	"github.com/huandu/facebook"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"server/credentials"
)

const datastore_key string = "user"

type AuthReq struct {
	Service string `json:"service"`
	Token   string `json:"token"`
}

type User struct {
	FbId string `datastore:"fb",json:"fb_id"`
}

var app = facebook.New(credentials.FACEBOOK_CLIENT_ID, credentials.FACEBOOK_CLIENT_SECRET)

func GetUserId(c context.Context, req AuthReq) (string, error) {
	switch req.Service {
	case "fb":
		log.Debugf(c, "getting user id from fb user")
		fbUserId, err := getFBUserId(c, req.Token)
		if err != nil {
			return "", err
		}
		log.Debugf(c, "fb responded, valid token with fbid: %s", fbUserId)
		return getOrCreateUserId(c, req.Service, fbUserId)
	}
	return "", errors.New("Unable to authenticate user")
}

func getOrCreateUserId(c context.Context, service string, id string) (string, error) {
	q := datastore.NewQuery(datastore_key).Filter(service+" =", id).KeysOnly().Limit(1)
	keys, err := q.GetAll(c, nil)
	if err != nil {
		return "", err
	}
	if len(keys) == 1 {
		log.Debugf(c, "queried for user, found key: %s", keys[0].Encode())
		return keys[0].Encode(), nil
	}

	// create a user
	var user User
	switch service {
	case "fb":
		user.FbId = id
	}
	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, datastore_key, nil), &user)
	log.Debugf(c, "queried for user, but didnt find. created new user: ", key.Encode())
	if err != nil {
		return "", err
	}

	return key.Encode(), nil
}

func getFBUserId(c context.Context, token string) (string, error) {
	session := app.Session(token)
	session.HttpClient = urlfetch.Client(c)

	res, err := session.Get("/me", nil)
	if err != nil {
		return "", err
	}

	var id string
	res.DecodeField("id", &id)

	return id, nil
}
