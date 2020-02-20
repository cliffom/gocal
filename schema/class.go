package schema

// Class defines a class
type Class struct {
	ID              string          `json:"id"`
	Type            ClassType       `json:"type"`
	Name            string          `json:"name"`
	Location        ClassFacility   `json:"location"`
	Room            Room            `json:"room"`
	Reservations    bool            `json:"reservations"`
	StartDatetime   string          `json:"start_datetime"`
	EndDate         string          `json:"end_date"`
	Recurrence      string          `json:"recurrence"`
	Instructor      ClassInstructor `json:"instructor"`
	Description     string          `json:"description"`
	Attendees       []ClassAttendee `json:"attendees"`
	TotalAttendance int             `json:"total_attendance"`
}
