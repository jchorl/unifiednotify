package service

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type Notification struct {
	Id      string `json:"id"`
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	Line3   string `json:"line3"`
	Date    int64  `json:"date"`
	URL     string `json:"url"`
	IconURL string `json:"iconUrl"`
}

type Service interface {
	GetNotifications(c context.Context, t *oauth2.Token)
}

func (n Notification) String() string {
	return fmt.Sprintf("Line1: %s\nLine2: %s\nLine3: %s\nDate: %d\nURL: %s\nIconURL: %s\n\n", n.Line1, n.Line2, n.Line3, n.Date, n.URL, n.IconURL)
}
