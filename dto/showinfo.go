package dto

import (
	"fmt"
	"log"
	"time"
)

type ShowInfo struct {
	Date  string
	Venue string
	Times string
}

func (show *ShowInfo) String() string {
	return fmt.Sprintf("Water Spots @ %s %s", show.Venue, show.Times)
}

func (showInfo *ShowInfo) ToScheduleEvent() ScheduleEvent {
	startDate, err := time.Parse("Mon, Jan 02 2006", fmt.Sprintf("%s %d", showInfo.Date, time.Now().Year()))
	if err != nil {
		log.Fatalf("Unable to parse date for event in correct format. got=%s", showInfo.Date)
	}

	// If the start date's month is less than the current month, then it's scheduled for next year.
	if startDate.Month() < time.Now().Month() {
		startDate.AddDate(1, 0, 0)
	}

	formattedDate := startDate.Format("2006-01-02")

	return ScheduleEvent{Summary: showInfo.String(), DateTime: formattedDate}
}
