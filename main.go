package main

import (
	"fmt"
	"log"

	"cerberus.com/ws-schedule-scraper/calendar"
	"cerberus.com/ws-schedule-scraper/config"
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
		fmt.Println(scheduledShow.String())
	}

	fmt.Println()
	fmt.Println("-------- GOOGLE CALENDAR EVENTS --------")

	googleCalendar := calendar.NewGoogleCalendar(&appConfig)
	events := googleCalendar.GetEvents()
	fmt.Println(events)

	// TODO: Merge events

	// TODO: Update google calendar with:
	// * new events
	// * updated events
}
