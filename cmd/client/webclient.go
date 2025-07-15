package client

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"cerberus.com/ws-schedule-scraper/config"
	"cerberus.com/ws-schedule-scraper/internal/dto"
)

type WebsiteClient struct {
	ScheduleUrl string
}

func NewWebsiteClient(config *config.AppConfig) (*WebsiteClient, error) {
	url, err := url.Parse(config.WebsiteScheduleUrl)
	if err != nil {
		return nil, err
	}

	return &WebsiteClient{ScheduleUrl: url.String()}, nil
}

func (wc *WebsiteClient) GetEvents() []dto.WebSiteEvent {
	response, err := http.Get(wc.ScheduleUrl)
	if err != nil {
		log.Fatalf("Error occurred getting JSON from %s: %v", wc.ScheduleUrl, err)
	}
	defer response.Body.Close()

	var webSiteEvents []dto.WebSiteEvent
	if err := json.NewDecoder(response.Body).Decode(&webSiteEvents); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return webSiteEvents
}
