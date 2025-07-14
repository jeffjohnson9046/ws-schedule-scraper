package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"cerberus.com/ws-schedule-scraper/config"
	"cerberus.com/ws-schedule-scraper/internal/dto"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleCalendarClient struct {
	CalendarId      string
	CredentialsFile string
	MaxResults      int64
	TimeZone        string
}

func NewGoogleCalendarClient(config *config.AppConfig) *GoogleCalendarClient {
	return &GoogleCalendarClient{CalendarId: config.CalendarId, CredentialsFile: config.CredentialsFile, MaxResults: config.MaxResults}
}

func (cal *GoogleCalendarClient) GetEvents() []dto.CalendarEvent {
	events := make([]dto.CalendarEvent, 0)
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

		events = append(events, dto.CalendarEvent{Summary: item.Summary, DateTime: date, EventId: item.Id})
	}

	return events
}

func (cal *GoogleCalendarClient) CreateEvents(websiteEvents []dto.WebSiteEvent) {
	newCalendarEvents := make([]dto.CalendarEvent, 0)
	for _, websiteEvent := range websiteEvents {
		newCalendarEvents = append(newCalendarEvents, dto.CalendarEvent{Summary: websiteEvent.String(), DateTime: websiteEvent.GetEventDateTime()})
	}

	ctx := context.Background()
	calendarService, err := calendar.NewService(ctx, option.WithCredentialsFile(cal.CredentialsFile))
	if err != nil {
		log.Fatalf("Could not get calendar service: %v", err)
	}

	for _, event := range newCalendarEvents {
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

func (cal *GoogleCalendarClient) UpdateEvents(eventsToUpdate []dto.CalendarEvent) {
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

func (cal *GoogleCalendarClient) DeleteEvents(eventsToDelete []dto.CalendarEvent) {
	ctx := context.Background()
	calendarService, err := calendar.NewService(ctx, option.WithCredentialsFile(cal.CredentialsFile))
	if err != nil {
		log.Fatalf("Could not get calendar service: %v", err)
	}

	for _, event := range eventsToDelete {
		fmt.Println(event)

		err := calendarService.Events.Delete(cal.CalendarId, event.EventId).Do()
		if err != nil {
			log.Fatalf("Unable to delete event %e", err)
		}

	}
}
