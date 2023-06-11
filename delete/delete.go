func deleteEvent(srv *calendar.Service, eventID string) error {
    err := srv.Events.Delete("primary", eventID).Do()
    if err != nil {
        return fmt.Errorf("failed to delete event: %v", err)
    }
    return nil
}
