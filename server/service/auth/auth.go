package auth

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"server/constants"
	"server/credentials"
)

var conf *oauth2.Config = &oauth2.Config{
	ClientID:     credentials.GOOGLE_CLIENT_ID,
	ClientSecret: credentials.GOOGLE_CLIENT_SECRET,
	RedirectURL:  constants.BASE_URL + "/googleoauthcallback",
	Scopes: []string{
		gmail.GmailReadonlyScope,
	},
	Endpoint: google.Endpoint,
}

func GetAuthUrl(userID string) string {
	return conf.AuthCodeURL(userID)
}

func GetToken(c context.Context, code string) (*oauth2.Token, error) {
	return conf.Exchange(c, code)
}

func GetConfig() *oauth2.Config {
	return conf
}
