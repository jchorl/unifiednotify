package gmail

import (
	"golang.org/x/net/context"
	"google.golang.org/api/gmail/v1"
	"html"
	"net/mail"
	"server/constants"
	"server/service"
	"server/service/auth"
	"server/tokenstore"
	"time"
)

type Message struct {
	Id           string
	Subject      string
	Snippet      string
	InternalDate time.Time
	Sender       string
}

func GetNotifications(c context.Context, t tokenstore.Token) ([]service.Notification, error) {
	client := auth.GetConfig(constants.GMAIL_SERVICE).Client(c, t.ToOauth())
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
	r, err := svc.Users.Messages.List("me").MaxResults(5).Q("in:inbox").Do()
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
		r, err := svc.Users.Messages.Get("me", msg.Id).Format("metadata").MetadataHeaders("From", "Subject").Fields("internalDate", "payload", "snippet").Do()
		if err != nil {
			return nil, err
		}
		msg.Snippet = html.UnescapeString(r.Snippet)
		// need to convert ms-epoch date to time.Time
		msg.InternalDate = time.Unix(r.InternalDate/1000, 0)
		for _, header := range r.Payload.Headers {
			if header.Name == "From" {
				msg.Sender = header.Value
			} else if header.Name == "Subject" {
				msg.Subject = header.Value
			}
		}
		parsed, err := mail.ParseAddress(msg.Sender)
		if err != nil {
			return nil, err
		}
		if parsed.Name != "" {
			msg.Sender = parsed.Name
		} else if parsed.Address != "" {
			msg.Sender = parsed.Address
		}
	}
	return messages, nil
}

func getNotificationsFromMessages(messages []*Message) []service.Notification {
	var notifications []service.Notification
	for _, msg := range messages {
		notifications = append(notifications, service.Notification{
			Id:      constants.GMAIL_SERVICE + msg.Id,
			Line1:   msg.Sender,
			Line2:   msg.Subject,
			Line3:   msg.Snippet,
			Date:    msg.InternalDate,
			URL:     "https://mail.google.com",
			IconURL: "https://trainerlearningcenter.withgoogle.com/assets/images/gmail.png",
		})
	}
	return notifications
}
