package main

import (
	"context"
	"fmt"
)

func handlerReset(state *State, cmd Command) error {
	contextBackground := context.Background()

	err := state.db.Reset(contextBackground)
	if err != nil {
		return err
	}

	fmt.Println("database reset successfully")

	return nil
}
