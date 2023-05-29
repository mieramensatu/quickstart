package read

import (
	"fmt"
	"log"

	"google.golang.org/api/calendar/v3"
)

func ReadCalendar(srv *calendar.Service, calendarId string, event *calendar.Event) {
	fetchedEvent, err := srv.Events.Get(calendarId, event.Id).Do()
	if err != nil {
		log.Fatalf("Unable to get event. %v\n", err)
	}
	fmt.Printf("Event fetched: %s\n", fetchedEvent.HtmlLink)
}