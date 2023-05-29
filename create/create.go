package create

import (
	"fmt"
	"time"

	"google.golang.org/api/calendar/v3"
)

func CreateEvent(srv *calendar.Service, summary string, description string, startTime time.Time, endTime time.Time) (*calendar.Event, error) {
	event := &calendar.Event{
		Summary:     summary,
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: startTime.Format(time.RFC3339),
			TimeZone: "Asia/Jakarta", // sesuai dengan zona waktu 
		},
		End: &calendar.EventDateTime{
			DateTime: endTime.Format(time.RFC3339),
			TimeZone: "Asia/Jakarta", // Sesuaikan dengan zona waktu
		},
	}

	event, err := srv.Events.Insert("primary", event).Do()
	if err != nil {
		return nil, fmt.Errorf("gagal membuat event: %v", err)
	}
	return event, nil
}
