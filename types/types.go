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

// NewUser creates a new user - v0.1.1 feature
func NewUser(id int, name string) User {
	return User{ID: id, Name: name}
}
