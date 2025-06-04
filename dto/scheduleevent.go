package dto

import (
	"fmt"
)

type ScheduleEvent struct {
	Summary  string
	DateTime string
	EventId  string
}

func (se *ScheduleEvent) String() {
	fmt.Printf("(id: %s) Summary: %s, DateTime: %s", se.EventId, se.Summary, se.DateTime)
}
