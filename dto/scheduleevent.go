package dto

import (
	"fmt"
)

type ScheduleEvent struct {
	Summary  string
	DateTime string
}

func (se *ScheduleEvent) String() {
	fmt.Printf("Summary: %s, DateTime: %s", se.Summary, se.DateTime)
}
