package read

import (
	"fmt"

	"google.golang.org/api/calendar/v3"
)

func getEvents(srv *calendar.Service) ([]*calendar.Event, error) {
	events, err := srv.Events.List("primary").Do()
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar event: %v", err)
	}
	return events.Items, nil
}

