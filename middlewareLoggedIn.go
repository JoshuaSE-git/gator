package main

import (
	"context"

	"github.com/JoshuaSE-git/gator/internal/database"
)

func middlewareLoggedIn(handler func(*State, Command, database.User) error) func(*State, Command) error {
	return func(state *State, cmd Command) error {
		ctx := context.Background()
		userName := state.cfg.CurrentUserName

		user, err := state.db.GetUser(ctx, userName)
		if err != nil {
			return err
		}

		err = handler(state, cmd, user)
		if err != nil {
			return err
		}

		return nil
	}
}
