package tokenstore

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"google.golang.org/appengine/datastore"
	"time"
)

const datastore_key string = "token"

type Token struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
	TokenType    string
	Kind         string
}

func (t Token) ToOauth() *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		Expiry:       t.Expiry,
		TokenType:    t.TokenType,
	}
}

func ToToken(t *oauth2.Token, kind string) *Token {
	return &Token{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		Expiry:       t.Expiry,
		TokenType:    t.TokenType,
		Kind:         kind,
	}
}

func GetTokensByUser(c context.Context, userId string) ([]Token, error) {
	key, err := datastore.DecodeKey(userId)
	if err != nil {
		return nil, err
	}
	q := datastore.NewQuery(datastore_key).Ancestor(key)

	var tokens []Token
	_, err = q.GetAll(c, &tokens)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func SaveToken(c context.Context, userId string, token *Token) error {
	key, err := datastore.DecodeKey(userId)
	if err != nil {
		return err
	}

	_, err = datastore.Put(c, datastore.NewIncompleteKey(c, datastore_key, key), token)
	return err
}
