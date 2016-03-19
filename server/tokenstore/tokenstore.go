package tokenstore

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"google.golang.org/appengine/datastore"
	"time"
)

type Token struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
	TokenType    string
}

func (t Token) toOauth() *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		Expiry:       t.Expiry,
		TokenType:    t.TokenType,
	}
}

func toToken(t *oauth2.Token) *Token {
	return &Token{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		Expiry:       t.Expiry,
		TokenType:    t.TokenType,
	}
}

func GetToken(c context.Context, userID string, kind string) (*oauth2.Token, error) {
	t := new(Token)
	err := datastore.Get(c, datastore.NewKey(c, kind, userID, 0, nil), t)
	if err != nil {
		return nil, err
	}
	return t.toOauth(), nil
}

func SaveToken(c context.Context, userID string, kind string, token *oauth2.Token) error {
	_, err := datastore.Put(c, datastore.NewKey(c, kind, userID, 0, nil), toToken(token))
	return err
}
