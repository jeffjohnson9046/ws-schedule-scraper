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

	websiteClient, err := client.NewWebsiteClient(appConfig)
	if err != nil {
		log.Fatalf("Error attempting to create website client: %v", err)
	}
	googleCalendarClient := client.NewGoogleCalendarClient(appConfig)

	sync.Execute(appConfig, websiteClient, googleCalendarClient)
}
