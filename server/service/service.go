package service

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type Notification struct {
	Line1   string
	Line2   string
	Line3   string
	Date    int64
	URL     string
	IconURL string
}

type Service interface {
	GetNotifications(c context.Context, t *oauth2.Token)
}

func (n Notification) String() string {
	return fmt.Sprintf("Line1: %s\nLine2: %s\nLine3: %s\nDate: %d\nURL: %s\nIconURL: %s\n\n", n.Line1, n.Line2, n.Line3, n.Date, n.URL, n.IconURL)
}
