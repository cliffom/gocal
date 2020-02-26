package schema_test

import (
	"testing"

	"github.com/cliffom/gocal/schema"
	"github.com/google/uuid"
	"syreclabs.com/go/faker"
)

func TestCreateListingFromClass(t *testing.T) {
	var attendees []schema.ClassAttendee
	attendee := schema.ClassAttendee{
		ID:   uuid.New().String(),
		Name: faker.Name().Name(),
	}
	attendees = append(attendees, attendee)

	class := schema.Class{
		Type: schema.ClassType{
			ID:   uuid.New().String(),
			Type: "Training",
		},
		Location: schema.ClassFacility{},
		Room: schema.Room{
			ID:       uuid.New().String(),
			Name:     "Weight Room",
			Capacity: faker.Number().NumberInt(2),
			Active:   true,
		},
		Reservations:  false,
		StartDatetime: "2020-02-06T15:00:00+0000",
		EndDate:       "2020-02-20",
		Recurrence:    "string",
		Instructor: schema.ClassInstructor{
			ID:   uuid.New().String(),
			Name: faker.Name().Name(),
		},
		Attendees:       attendees,
		TotalAttendance: faker.Number().NumberInt(2),
	}

	_, err := class.CreateListing()

	if err != nil {
		t.Errorf("Unable to create ClassListing from Class")
	}
}
