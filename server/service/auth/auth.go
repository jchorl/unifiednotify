package auth

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"server/constants"
	"server/credentials"
)

func GetAuthUrl(userID string, service string) string {
	conf := GetConfig(service)
	return conf.AuthCodeURL(userID)
}

func GetToken(c context.Context, code string, service string) (*oauth2.Token, error) {
	conf := GetConfig(service)
	return conf.Exchange(c, code)
}

func GetConfig(service string) *oauth2.Config {
	var conf *oauth2.Config
	switch service {
	case constants.GMAIL_SERVICE:
		conf = &oauth2.Config{
			ClientID:     credentials.GMAIL_CLIENT_ID,
			ClientSecret: credentials.GMAIL_CLIENT_SECRET,
			RedirectURL:  constants.BASE_URL + "/auth/gmail/callback",
			Scopes: []string{
				gmail.GmailReadonlyScope,
			},
			Endpoint: google.Endpoint,
		}
	}
	return conf
}
