package sync

import (
	"cerberus.com/ws-schedule-scraper/cmd/client"
	"cerberus.com/ws-schedule-scraper/config"
	"cerberus.com/ws-schedule-scraper/internal/dto"
)

type CalendarClient interface {
	GetEvents() []dto.CalendarEvent
	CreateEvents([]dto.WebSiteEvent)
	DeleteEvents([]dto.CalendarEvent)
	UpdateEvents([]dto.CalendarEvent)
}

func Execute(config *config.AppConfig, calendar CalendarClient) {
	websiteEvents := client.GetEvents(config)
	calendarEvents := calendar.GetEvents()

	// TEST/DEBUG ----------------------------------------------------
	// fmt.Println("-------- WATER SPOTS WEBSITE EVENTS --------")
	// for _, wsEvent := range websiteEvents {
	// 	fmt.Println(wsEvent.String())
	// }

	// fmt.Println()
	// fmt.Println("-------- GOOGLE CALENDAR EVENTS --------")
	// for _, calEvent := range calendarEvents {
	// 	fmt.Println(calEvent.String())
	// }
	// END TEST/DEBUG ------------------------------------------------

	calendarEventsByDate := make(map[string]dto.CalendarEvent)
	for _, calendarEvent := range calendarEvents {
		calendarEventsByDate[calendarEvent.DateTime] = calendarEvent
	}

	websiteEventsByDate := make(map[string]dto.WebSiteEvent)
	for _, websiteEvent := range websiteEvents {
		websiteEventsByDate[websiteEvent.GetEventDateTime()] = websiteEvent
	}

	eventsToDelete := make([]dto.CalendarEvent, 0)
	eventsToCreate := make([]dto.WebSiteEvent, 0)
	eventsToUpdate := make([]dto.CalendarEvent, 0)

	for _, calendarEvent := range calendarEvents {
		if _, found := websiteEventsByDate[calendarEvent.DateTime]; !found {
			eventsToDelete = append(eventsToDelete, calendarEvent)
		}
	}

	for _, websiteEvent := range websiteEvents {
		if existingCalendarEvent, found := calendarEventsByDate[websiteEvent.GetEventDateTime()]; found {
			if existingCalendarEvent.Summary != websiteEvent.String() {
				existingCalendarEvent.Summary = websiteEvent.String()

				eventsToUpdate = append(eventsToUpdate, existingCalendarEvent)
			}
		} else {
			eventsToCreate = append(eventsToCreate, websiteEvent)
		}
	}

	calendar.DeleteEvents(eventsToDelete)

	calendar.CreateEvents(eventsToCreate)

	calendar.UpdateEvents(eventsToUpdate)
}
