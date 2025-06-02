package calendar

import (
	"context"
	"log"
	"time"

	"cerberus.com/ws-schedule-scraper/config"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleCalendar struct {
	CalendarId      string
	CredentialsFile string
	MaxResults      int64
}

func NewGoogleCalendar(config *config.AppConfig) *GoogleCalendar {
	return &GoogleCalendar{CalendarId: config.CalendarId, CredentialsFile: config.CredentialsFile, MaxResults: config.MaxResults}
}

func (cal *GoogleCalendar) GetEvents() []Event {
	events := make([]Event, 0)
	ctx := context.Background()

	calendarService, err := calendar.NewService(ctx, option.WithCredentialsFile(cal.CredentialsFile))
	if err != nil {
		log.Fatalf("Could not get calendar service: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	calendarEvents, err := calendarService.Events.List(cal.CalendarId).
		SingleEvents(true).
		TimeMin(t).
		MaxResults(cal.MaxResults).
		OrderBy("startTime").
		Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}

	for _, item := range calendarEvents.Items {
		date := item.Start.DateTime
		if date == "" {
			date = item.Start.Date
		}

		events = append(events, Event{Summary: item.Summary, DateTime: date})
	}

	return events
}
