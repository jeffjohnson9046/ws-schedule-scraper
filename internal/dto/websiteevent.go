package dto

import (
	"fmt"
	"log"
	"time"
)

type WebSiteEvent struct {
	Date          string `json:"Date"`
	Time          string `json:"Time"`
	Venue         string `json:"Venue"`
	City          string `json:"City"`
	State         string `json:"State"`
	MapUrl        string `json:"MapURL"`
	OtherBands    string `json:"OtherBands"`
	PhotoFileName string `json:"PhotoFilename"`
	PromoBlurb    string `json:"PromoBlurb"`
	ShowTitle     string `json:"ShowTitle"`
}

func (wse *WebSiteEvent) String() string {
	return fmt.Sprintf("Water Spots @ %s %s", wse.Venue, wse.Time)
}

func (wse *WebSiteEvent) GetEventDateTime() string {
	startDate, err := time.Parse("Mon, Jan 02 2006", fmt.Sprintf("%s %d", wse.Date, time.Now().Year()))
	if err != nil {
		log.Fatalf("Unable to parse date for event in correct format. got=%s", wse.Date)
	}

	// If the start date's month is less than the current month, then it's scheduled for next year.
	if startDate.Month() < time.Now().Month() {
		startDate.AddDate(1, 0, 0)
	}

	return startDate.Format("2006-01-02")
}
