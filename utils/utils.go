package utils

import "github.com/geyslan/go-modules-test/types"

// ProcessUser processes a user
func ProcessUser(user types.User) string {
	return "Processing user: " + user.Name
}

// ProcessEvent processes an event
func ProcessEvent(event types.Event) string {
	return "Processing event: " + event.Type
}

// CreateAndProcessUser creates and processes a user - v0.1.1 feature
func CreateAndProcessUser(id int, name string) string {
	user := types.NewUser(id, name)
	return "Created and " + ProcessUser(user)
}
