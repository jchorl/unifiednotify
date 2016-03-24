package constants

import (
	"google.golang.org/appengine"
)

const GMAIL_SERVICE string = "gmail"
const GCAL_SERVICE string = "gcal"

var BASE_URL string

func init() {
	if appengine.IsDevAppServer() {
		BASE_URL = "http://localhost:8080"
	} else {
		BASE_URL = "https://unifiednotify.appspot.com"
	}
}
