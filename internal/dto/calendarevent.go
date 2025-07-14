package dto

import (
	"fmt"
)

type CalendarEvent struct {
	Summary  string
	DateTime string
	EventId  string
}

func (ce *CalendarEvent) String() string {
	return fmt.Sprintf("(id: %s) Summary: %s, DateTime: %s", ce.EventId, ce.Summary, ce.DateTime)
}
