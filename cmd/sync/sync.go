package sync

import (
	"fmt"

	"cerberus.com/ws-schedule-scraper/config"
	"cerberus.com/ws-schedule-scraper/internal/dto"
)

type WebsiteClient interface {
	GetEvents() []dto.WebSiteEvent
}

type CalendarClient interface {
	GetEvents() []dto.CalendarEvent
	CreateEvents([]dto.WebSiteEvent)
	DeleteEvents([]dto.CalendarEvent)
	UpdateEvents([]dto.CalendarEvent)
}

func Execute(config *config.AppConfig, websiteClient WebsiteClient, calendarClient CalendarClient) {
	websiteEvents := websiteClient.GetEvents()
	calendarEvents := calendarClient.GetEvents()

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
		websiteEventsByDate[websiteEvent.Date] = websiteEvent
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
		if existingCalendarEvent, found := calendarEventsByDate[websiteEvent.Date]; found {
			webSiteEventSummary := websiteEvent.String()
			if existingCalendarEvent.Summary != webSiteEventSummary {
				existingCalendarEvent.Summary = webSiteEventSummary

				eventsToUpdate = append(eventsToUpdate, existingCalendarEvent)
			}
		} else {
			eventsToCreate = append(eventsToCreate, websiteEvent)
		}
	}

	fmt.Println("-------- RESULTS --------")
	fmt.Println("--- DELETES ---")
	for _, e := range eventsToDelete {
		fmt.Println(e.String())
	}
	fmt.Println()

	fmt.Println("--- CREATES ---")
	for _, e := range eventsToCreate {
		fmt.Println(e.String())
	}
	fmt.Println()

	fmt.Println("--- UPDATES ---")
	for _, e := range eventsToUpdate {
		fmt.Println(e.String())
	}
	fmt.Println()

	calendarClient.DeleteEvents(eventsToDelete)

	calendarClient.CreateEvents(eventsToCreate)

	calendarClient.UpdateEvents(eventsToUpdate)
}
