package gcal

import (
	"golang.org/x/net/context"
	"google.golang.org/api/calendar/v3"
	"server/constants"
	"server/service"
	"server/service/auth"
	"server/tokenstore"
	"time"
)

type Calendar struct {
	Id      string
	Summary string
}

type Event struct {
	Id          string
	Link        string
	Start       time.Time
	AllDay      bool
	Summary     string
	Description string
	CalSummary  string
}

func GetNotifications(c context.Context, t tokenstore.Token) ([]service.Notification, error) {
	client := auth.GetConfig(constants.GCAL_SERVICE).Client(c, t.ToOauth())
	svc, err := calendar.New(client)
	if err != nil {
		return nil, err
	}
	calendars, err := getCalendars(svc)
	if err != nil {
		return nil, err
	}
	events, err := populateAndAggregate(svc, calendars)
	if err != nil {
		return nil, err
	}
	notifications := getNotificationsFromEvents(events)
	return notifications, nil
}

func getCalendars(svc *calendar.Service) ([]*Calendar, error) {
	r, err := svc.CalendarList.List().Fields("items/id", "items/summary").Do()
	if err != nil {
		return nil, err
	}
	var calendars []*Calendar
	for _, cal := range r.Items {
		calendars = append(calendars, &Calendar{
			Id:      cal.Id,
			Summary: cal.Summary,
		})
	}
	return calendars, nil
}

func populateAndAggregate(svc *calendar.Service, calendars []*Calendar) ([]*Event, error) {
	var events []*Event
	// get start and end time
	now := time.Now()
	oneDay := now.Add(time.Hour * 24)
	for _, cal := range calendars {
		r, err := svc.Events.List(cal.Id).MaxResults(5).OrderBy("startTime").SingleEvents(true).TimeMin(now.Format(time.RFC3339)).TimeMax(oneDay.Format(time.RFC3339)).Fields("items/id", "items/description", "items/htmlLink", "items/start", "items/summary").Do()
		if err != nil {
			return nil, err
		}
		for _, ev := range r.Items {
			// try to parse the time
			// this could be an all day event
			start, allDay, err := parseStart(ev.Start)
			if err != nil {
				return nil, err
			}
			events = append(events, &Event{
				Id:          ev.Id,
				Link:        ev.HtmlLink,
				Start:       start,
				AllDay:      allDay,
				Summary:     ev.Summary,
				Description: ev.Description,
				CalSummary:  cal.Summary,
			})
		}
	}
	return events, nil
}

func parseStart(tm *calendar.EventDateTime) (time.Time, bool, error) {
	if tm.DateTime != "" {
		parsed, err := time.Parse(time.RFC3339, tm.DateTime)
		return parsed, false, err
	}
	parsed, err := time.Parse("2006-01-02", tm.Date)
	// set for 11:59:59:...
	parsed.Add(time.Duration(24)*time.Hour - time.Duration(1)*time.Second)
	return parsed, false, err
}

func getNotificationsFromEvents(events []*Event) []service.Notification {
	var notifications []service.Notification
	for _, ev := range events {
		notifications = append(notifications, service.Notification{
			Id:      constants.GCAL_SERVICE + ev.Id,
			Line1:   ev.Summary,
			Line2:   ev.CalSummary,
			Line3:   ev.Description,
			Date:    ev.Start,
			URL:     ev.Link,
			IconURL: "http://icons.iconarchive.com/icons/dtafalonso/android-lollipop/128/calendar-icon.png",
		})
	}
	return notifications
}
