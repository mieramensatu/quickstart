package delete

import (
	"fmt"
	"log"

	"google.golang.org/api/calendar/v3"
)

func DeleteCalendar(srv *calendar.Service, calendarId string, event *calendar.Event) {
	err := srv.Events.Delete(calendarId, event.Id).Do()
	if err != nil {
		log.Fatalf("Unable to delete event. %v\n", err)
	}
	fmt.Println("Event deleted")
}
