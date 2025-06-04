package main

import (
	"fmt"
	"log"

	"cerberus.com/ws-schedule-scraper/calendar"
	"cerberus.com/ws-schedule-scraper/config"
	"cerberus.com/ws-schedule-scraper/dto"
	"cerberus.com/ws-schedule-scraper/scraper"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %e", err)
	}

	appConfig := config.AppConfig{}

	err = env.Parse(&appConfig)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	scraper := scraper.New(&appConfig)
	scheduledShowResults := scraper.Scrape()

	fmt.Println("-------- WATER SPOTS WEBSITE EVENTS --------")

	for _, scheduledShow := range scheduledShowResults {
		fmt.Println(scheduledShow.ToScheduleEvent())
	}

	fmt.Println()
	fmt.Println("-------- GOOGLE CALENDAR EVENTS --------")

	googleCalendar := calendar.NewGoogleCalendar(&appConfig)
	calendarEvents := googleCalendar.GetEvents()

	eventsOnCalendar := make(map[string]dto.ScheduleEvent)
	for _, calendarEvent := range calendarEvents {
		eventsOnCalendar[calendarEvent.DateTime] = calendarEvent
	}

	eventsToCreate := make([]dto.ShowInfo, 0)
	eventsToUpdate := make([]dto.ScheduleEvent, 0)
	for _, scheduledShowResult := range scheduledShowResults {
		resultAsEvent := scheduledShowResult.ToScheduleEvent()

		if existingCalendarEvent, ok := eventsOnCalendar[resultAsEvent.DateTime]; ok {
			if existingCalendarEvent.Summary != scheduledShowResult.String() {
				existingCalendarEvent.Summary = scheduledShowResult.String()

				eventsToUpdate = append(eventsToUpdate, existingCalendarEvent)
			}
		} else {
			eventsToCreate = append(eventsToCreate, scheduledShowResult)
		}
	}

	// Create new events
	googleCalendar.CreateEvents(eventsToCreate)

	// Update existing events
	googleCalendar.UpdateEvents(eventsToUpdate)
}
