package service

import (
	"fmt"
	"golang.org/x/net/context"
	"server/tokenstore"
	"time"
)

var SERVICE_MAPPINGS map[string]func(c context.Context, t tokenstore.Token) ([]Notification, error) = map[string]func(c context.Context, t tokenstore.Token) ([]Notification, error){}

type Notification struct {
	Id      string    `json:"id"`
	Line1   string    `json:"line1"`
	Line2   string    `json:"line2"`
	Line3   string    `json:"line3"`
	Date    time.Time `json:"date"`
	URL     string    `json:"url"`
	IconURL string    `json:"iconUrl"`
}

type ByDate []Notification

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date.After(a[j].Date) }

func (n Notification) String() string {
	return fmt.Sprintf("Line1: %s\nLine2: %s\nLine3: %s\nDate: %d\nURL: %s\nIconURL: %s\n\n", n.Line1, n.Line2, n.Line3, n.Date, n.URL, n.IconURL)
}
