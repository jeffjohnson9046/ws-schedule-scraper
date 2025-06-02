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
	return fmt.Sprintf("The Water Spots at %s %s", show.Venue, show.Times)
}

func (showInfo *ShowInfo) ToScheduleEvent() ScheduleEvent {
	startDate, err := time.Parse("Mon, Jan 02 2006", fmt.Sprintf("%s %d", showInfo.Date, time.Now().Year()))
	if err != nil {
		log.Fatalf("Unable to parse date for event in correct format. got=%s", showInfo.Date)
	}

	// TODO: pretty sure startDate.Format(-- the pattern --) should work here, but startDate.format("2006-01-01") kept giving back the same date, regardless of input
	formattedDate := fmt.Sprintf("%d-%02d-%02d", startDate.Year(), startDate.Month(), startDate.Day())

	return ScheduleEvent{Summary: showInfo.String(), DateTime: formattedDate}
}
