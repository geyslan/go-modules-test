package types

// User represents a user in the system
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Event represents an event
type Event struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
