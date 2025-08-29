package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JoshuaSE-git/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(state *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username required for registration")
	}

	id := uuid.New()
	currentTime := time.Now()
	name := cmd.args[0]
	params := database.CreateUserParams{
		ID:        id,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      name,
	}
	contextBackground := context.Background()

	user, err := state.db.GetUser(contextBackground, name)
	if err == nil {
		return fmt.Errorf("user %q already exists", user.Name)
	}

	user, err = state.db.CreateUser(contextBackground, params)
	if err != nil {
		return err
	}

	state.cfg.SetUser(user.Name)
	fmt.Printf("%q was successfully registered\n", user.Name)

	return nil
}
