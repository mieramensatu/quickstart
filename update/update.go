package update

import (
	"fmt"

	"google.golang.org/api/calendar/v3"
)

func updateEvent(srv *calendar.Service, event *calendar.Event) (*calendar.Event, error) {
	event, err := srv.Events.Update("primary", event.Id, event).Do()
	if err != nil {
		return nil, fmt.Errorf("gagal mengupdate event: %v", err)
	}
	return event, nil
}
