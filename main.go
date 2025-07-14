package main

import (
	"log"

	"cerberus.com/ws-schedule-scraper/cmd/client"
	"cerberus.com/ws-schedule-scraper/cmd/sync"
	"cerberus.com/ws-schedule-scraper/config"
)

func main() {
	appConfig, err := config.NewAppConfig()
	if err != nil {
		log.Fatalf("Error loading app configuration: %v", err)
	}

	googleCalendar := client.NewGoogleCalendarClient(appConfig)

	sync.Execute(appConfig, googleCalendar)
}
