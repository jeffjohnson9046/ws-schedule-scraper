package client

import (
	"encoding/json"
	"log"
	"net/http"

	"cerberus.com/ws-schedule-scraper/config"
	"cerberus.com/ws-schedule-scraper/internal/dto"
)

func GetEvents(config *config.AppConfig) []dto.WebSiteEvent {
	response, err := http.Get(config.WebsiteScheduleUrl)
	if err != nil {
		log.Fatalf("Error occurred getting JSON from %s: %v", config.WebsiteScheduleUrl, err)
	}
	defer response.Body.Close()

	var webSiteEvents []dto.WebSiteEvent
	if err := json.NewDecoder(response.Body).Decode(&webSiteEvents); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return webSiteEvents
}
