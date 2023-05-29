package delete

import (
	"fmt"

	"google.golang.org/api/calendar/v3"
)

func deleteEvent(srv *calendar.Service, eventID string) error {
	err := srv.Events.Delete("primary", eventID).Do()
	if err != nil {
		return fmt.Errorf("gagal menghapus event: %v", err)
	}
	return nil
}
