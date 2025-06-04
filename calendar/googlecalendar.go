package calendar

import (
	"context"
	"fmt"
	"log"
	"time"

	"cerberus.com/ws-schedule-scraper/config"
	"cerberus.com/ws-schedule-scraper/dto"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleCalendar struct {
	CalendarId      string
	CredentialsFile string
	MaxResults      int64
	TimeZone        string
}

func NewGoogleCalendar(config *config.AppConfig) *GoogleCalendar {
	return &GoogleCalendar{CalendarId: config.CalendarId, CredentialsFile: config.CredentialsFile, MaxResults: config.MaxResults}
}

func (cal *GoogleCalendar) GetEvents() []dto.ScheduleEvent {
	events := make([]dto.ScheduleEvent, 0)
	ctx := context.Background()

	calendarService, err := calendar.NewService(ctx, option.WithCredentialsFile(cal.CredentialsFile))
	if err != nil {
		log.Fatalf("Could not get calendar service: %v", err)
	}

	minTime := time.Now().Format(time.RFC3339)
	calendarEvents, err := calendarService.Events.List(cal.CalendarId).
		SingleEvents(true).
		TimeMin(minTime).
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

		events = append(events, dto.ScheduleEvent{Summary: item.Summary, DateTime: date})
	}

	return events
}

func (cal *GoogleCalendar) CreateEvents(showInfos []dto.ShowInfo) {
	eventsToSchedule := make([]dto.ScheduleEvent, 0)
	for _, show := range showInfos {
		eventsToSchedule = append(eventsToSchedule, show.ToScheduleEvent())
	}

	ctx := context.Background()
	calendarService, err := calendar.NewService(ctx, option.WithCredentialsFile(cal.CredentialsFile))
	if err != nil {
		log.Fatalf("Could not get calendar service: %v", err)
	}

	for _, event := range eventsToSchedule {
		fmt.Println(event)
		eventDate := calendar.EventDateTime{Date: event.DateTime, TimeZone: cal.TimeZone}

		newEvent := calendar.Event{Summary: event.Summary, Start: &eventDate, End: &eventDate}

		result, err := calendarService.Events.Insert(cal.CalendarId, &newEvent).Do()
		if err != nil {
			log.Fatalf("Unable to create event: %e", err)
		}

		fmt.Printf(">> Created event: %s (id: %s)\n", result.Summary, result.Id)
	}
}

func (cal *GoogleCalendar) UpdateEvents(eventsToUpdate []dto.ScheduleEvent) {
	ctx := context.Background()
	calendarService, err := calendar.NewService(ctx, option.WithCredentialsFile(cal.CredentialsFile))
	if err != nil {
		log.Fatalf("Could not get calendar service: %v", err)
	}

	for _, event := range eventsToUpdate {
		fmt.Println(event)
		eventDate := calendar.EventDateTime{Date: event.DateTime, TimeZone: cal.TimeZone}

		eventToUpdate := calendar.Event{Id: event.EventId, Summary: event.Summary, Start: &eventDate, End: &eventDate}

		result, err := calendarService.Events.Update(cal.CalendarId, eventToUpdate.Id, &eventToUpdate).Do()
		if err != nil {
			log.Fatalf("Unable to create event: %e", err)
		}

		fmt.Printf(">> Created event: %s (id: %s)\n", result.Summary, result.Id)
	}
}
