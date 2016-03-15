package gmail

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"server/service"
	"server/service/auth"
)

type Message struct {
	Id           string
	Subject      string
	Snippet      string
	InternalDate int64
	Sender       string
}

func GetNotifications(c context.Context, t *oauth2.Token) ([]service.Notification, error) {
	client := auth.GetConfig().Client(c, t)
	svc, err := gmail.New(client)
	if err != nil {
		return nil, err
	}
	messages, err := getIncompleteMessages(svc)
	if err != nil {
		return nil, err
	}
	messages, err = populateMessages(svc, messages)
	if err != nil {
		return nil, err
	}
	notifications := getNotificationsFromMessages(messages)
	return notifications, nil
}

func getIncompleteMessages(svc *gmail.Service) ([]*Message, error) {
	req := svc.Users.Messages.List("me").MaxResults(10).Q("in:inbox")
	r, err := req.Do()
	if err != nil {
		return nil, err
	}
	var messages []*Message
	for _, msg := range r.Messages {
		messages = append(messages, &Message{
			Id: msg.Id,
		})
	}
	return messages, nil
}

func populateMessages(svc *gmail.Service, messages []*Message) ([]*Message, error) {
	for _, msg := range messages {
		req := svc.Users.Messages.Get("me", msg.Id).Format("metadata").MetadataHeaders("from", "subject").Fields("internalDate", "payload", "snippet")
		r, err := req.Do()
		if err != nil {
			return nil, err
		}
		msg.Snippet = r.Snippet
		msg.InternalDate = r.InternalDate
		msg.Subject = r.Payload.Headers[0].Value
		msg.Subject = r.Payload.Headers[1].Value
	}
	return messages, nil
}

func getNotificationsFromMessages(messages []*Message) []service.Notification {
	var notifications []service.Notification
	for _, msg := range messages {
		notifications = append(notifications, service.Notification{
			Line1:   msg.Sender,
			Line2:   msg.Subject,
			Line3:   msg.Snippet,
			Date:    msg.InternalDate,
			URL:     "https://gmail.com",
			IconURL: "https://trainerlearningcenter.withgoogle.com/assets/images/gmail.png",
		})
	}
	return notifications
}
