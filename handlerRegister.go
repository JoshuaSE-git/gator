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
	context := context.Background()

	params := database.CreateUserParams{
		ID:        id,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      name,
	}

	_, err := state.db.GetUser(context, name)
	if err == nil {
		return fmt.Errorf("user %q already exists", name)
	}

	_, err = state.db.CreateUser(context, params)
	if err != nil {
		return err
	}

	state.cfg.SetUser(name)

	fmt.Printf("%q was successfully registered\n", name)

	return nil
}
