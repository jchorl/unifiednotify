package constants

import (
	"google.golang.org/appengine"
)

var BASE_URL string

const GMAIL_SERVICE string = "gmail"

func init() {
	if appengine.IsDevAppServer() {
		BASE_URL = "http://localhost:8080"
	} else {
		BASE_URL = "https://unifiednotify.appspot.com"
	}
}
