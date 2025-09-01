package main

import (
	"context"
	"fmt"
)

func handlerReset(state *State, cmd Command) error {
	ctx := context.Background()

	err := state.db.Reset(ctx)
	if err != nil {
		return err
	}

	err = state.db.ResetFeeds(ctx)
	if err != nil {
		return err
	}

	fmt.Println("database reset successfully")

	return nil
}
