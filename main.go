package main

import (
	"fmt"
	"github.com/geyslan/go-modules-test/types"
	"github.com/geyslan/go-modules-test/utils"
)

func main() {
	user := types.User{ID: 1, Name: "Test User"}
	event := types.Event{Type: "login", Data: "user logged in"}
	
	fmt.Println(utils.ProcessUser(user))
	fmt.Println(utils.ProcessEvent(event))
}
