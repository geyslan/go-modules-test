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
