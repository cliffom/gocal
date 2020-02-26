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

// CreateListing creates and returns a ClassListing
func (c Class) CreateListing() (*ClassListing, *Error) {
	if c.StartDatetime == "" {
		err := &Error{
			Message: "CreateListing Failed",
			Detail:  "StartDatetime is required",
		}

		return nil, err
	}

	listing := &ClassListing{
		ID:              "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		Type:            c.Type,
		Name:            "Body sculpt",
		Location:        c.Location,
		Room:            c.Room,
		StartDatetime:   c.StartDatetime,
		EndDatetime:     c.EndDate,
		Instructor:      c.Instructor.Name,
		TotalAttendance: 5,
		TotalRegistered: 10,
	}

	return listing, nil
}
