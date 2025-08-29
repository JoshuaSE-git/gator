package main

import (
	"context"
	"fmt"
)

func handlerUsers(state *State, cmd Command) error {
	contextBackground := context.Background()

	users, err := state.db.GetUsers(contextBackground)
	if err != nil {
		return err
	}

	if len(users) == 0 {
		fmt.Println("no users registered")
		return nil
	}

	for _, user := range users {
		if user.Name == state.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}

	return nil
}
