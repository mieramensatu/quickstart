package update

import (
	"fmt"
	"log"

	"google.golang.org/api/calendar/v3"
)

func UpdateCalendar(srv *calendar.Service, calendarId string, event *calendar.Event) {
	updatedEvent, err := srv.Events.Update(calendarId, event.Id, event).Do()
	if err != nil {
		log.Fatalf("Unable to update event. %v\n", err)
	}
	fmt.Printf("Event updated: %s\n", updatedEvent.HtmlLink)
}