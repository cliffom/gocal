package schema

// Error defines a message body to return when things go wrong
type Error struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
