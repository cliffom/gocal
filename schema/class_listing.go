package schema

// ClassListing defines a created class
type ClassListing struct {
	ID              string        `json:"id"`
	Type            ClassType     `json:"type"`
	Name            string        `json:"name"`
	Location        ClassFacility `json:"location"`
	Room            Room          `json:"room"`
	StartDatetime   string        `json:"start_datetime"`
	EndDatetime     string        `json:"end_datetime"`
	Instructor      string        `json:"instructor"`
	TotalAttendance int           `json:"total_attendance"`
	TotalRegistered int           `json:"total_registered"`
}
