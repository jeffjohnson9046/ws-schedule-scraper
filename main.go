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
		log.Fatal("Error loading .env file")
	}

	cfg := config.AppConfig{}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	scraper := scraper.New(&cfg)
	showResults := scraper.Scrape()

	fmt.Println("-------- WATER SPOTS WEBSITE EVENTS --------")

	for _, show := range showResults {
		fmt.Println(show.String())
	}

	fmt.Println()
	fmt.Println("-------- GOOGLE CALENDAR EVENTS --------")

	cal := calendar.New(&cfg)
	cal.GetEvents()

	// TODO: Merge events

	// TODO: Update google calendar with:
	// * new events
	// * updated events
}
