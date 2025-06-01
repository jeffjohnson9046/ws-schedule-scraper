package calendar

import (
	"context"
	"fmt"
	"log"
	"time"

	"cerberus.com/ws-schedule-scraper/config"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func GetEvents(config config.AppConfig) {
	ctx := context.Background()

	calendarService, err := calendar.NewService(ctx, option.WithCredentialsFile(config.CredentialsFile))
	if err != nil {
		log.Fatalf("Could not get calendar service: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := calendarService.Events.List(config.CalendarId).
		SingleEvents(true).
		TimeMin(t).
		MaxResults(config.MaxResults).
		OrderBy("startTime").
		Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}

	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}

			parsedTime, err := time.Parse(time.RFC3339, date)
			if err != nil {
				log.Fatalf("could not parse time. got=%s", date)
			}

			fmt.Printf("%v (%v), parsed: %v\n", item.Summary, date, parsedTime)
		}
	}
}
