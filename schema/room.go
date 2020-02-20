package schema

// Room defines a room
type Room struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Active   bool   `json:"active"`
}
