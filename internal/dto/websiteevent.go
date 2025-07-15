package dto

import (
	"fmt"
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
	return fmt.Sprintf("Water Spots @ %s (%s)", wse.Venue, wse.Time)
}
