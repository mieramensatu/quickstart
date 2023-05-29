package delete

import (
	"google.golang.org/api/calendar/v3"
	"log"

	"github.com/mieramensatu/quickstart.git/create"
)

// DeleteEvent deletes an existing event from Google Calendar
func DeleteEvent(srv *calendar.Service, eventID string) error {
	err := srv.Events.Delete("primary", eventID).Do()
	if err != nil {
		log.Fatalf("Unable to delete event: %v\n", err)
		return err
	}
	log.Printf("Event deleted successfully.\n")
	return nil
}
