package schema

// ClassAttendee defines an attendee of a class
type ClassAttendee struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CheckedIn bool   `json:"checked_in"`
}
