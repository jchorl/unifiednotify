package constants

import (
	"google.golang.org/appengine"
)

const GMAIL_SERVICE string = "gmail"

var BASE_URL string

func init() {
	if appengine.IsDevAppServer() {
		BASE_URL = "http://localhost:8080"
	} else {
		BASE_URL = "https://unifiednotify.appspot.com"
	}
}
