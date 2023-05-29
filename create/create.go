package create

import (
	"fmt"
	"log"

	"google.golang.org/api/calendar/v3"
)

func CreateCalendar(srv *calendar.Service, calendarId string, event *calendar.Event) {
	newEvent, err := srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", newEvent.HtmlLink)
}