package crud

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	// "golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// CalendarEvent represents a calendar event.
type CalendarEvent struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}

// Retrieve a token, saves the token, then returns the generated client.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, fmt.Errorf("unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve token from web: %v", err)
	}

	return tok, nil
}

// Request a token from the web, then returns the retrieved token.
func GetTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok

}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// CreateEvent creates a new calendar event.
func CreateEvent(client *http.Client, calendarID string, event *CalendarEvent) error {
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("unable to retrieve Calendar client: %v", err)
	}

	eventStartTime, err := time.Parse(time.RFC3339, event.StartTime)
	if err != nil {
		return fmt.Errorf("invalid start time: %v", err)
	}

	eventEndTime, err := time.Parse(time.RFC3339, event.EndTime)
	if err != nil {
		return fmt.Errorf("invalid end time: %v", err)
	}

	newEvent := &calendar.Event{
		Summary:     event.Summary,
		Description: event.Description,
		Start: &calendar.EventDateTime{
			DateTime: eventStartTime.Format(time.RFC3339),
		},
		End: &calendar.EventDateTime{
			DateTime: eventEndTime.Format(time.RFC3339),
		},
	}

	_, err = srv.Events.Insert(calendarID, newEvent).Do()
	if err != nil {
		return fmt.Errorf("unable to create event: %v", err)
	}

	return nil
}

// GetEvents retrieves a list of calendar events.
func GetEvents(client *http.Client, calendarID string, maxResults int) ([]*calendar.Event, error) {
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List(calendarID).ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(int64(maxResults)).OrderBy("startTime").Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve events: %v", err)
	}

	return events.Items, nil

}

// UpdateEvent updates an existing calendar event.
func UpdateEvent(client *http.Client, calendarID, eventID string, event *CalendarEvent) error {
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("unable to retrieve Calendar client: %v", err)
	}

	eventStartTime, err := time.Parse(time.RFC3339, event.StartTime)
	if err != nil {
		return fmt.Errorf("invalid start time: %v", err)
	}

	eventEndTime, err := time.Parse(time.RFC3339, event.EndTime)
	if err != nil {
		return fmt.Errorf("invalid end time: %v", err)
	}

	updatedEvent := &calendar.Event{
		Summary:     event.Summary,
		Description: event.Description,
		Start: &calendar.EventDateTime{
			DateTime: eventStartTime.Format(time.RFC3339),
		},
		End: &calendar.EventDateTime{
			DateTime: eventEndTime.Format(time.RFC3339),
		},
	}

	_, err = srv.Events.Update(calendarID, eventID, updatedEvent).Do()
	if err != nil {
		return fmt.Errorf("unable to update event: %v", err)
	}

	return nil
}

// DeleteEvent deletes a calendar event.
func DeleteEvent(client *http.Client, calendarID, eventID string) error {
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("unable to retrieve Calendar client: %v", err)
	}

	err = srv.Events.Delete(calendarID, eventID).Do()
	if err != nil {
		return fmt.Errorf("unable to delete event: %v", err)
	}

	return nil

}
