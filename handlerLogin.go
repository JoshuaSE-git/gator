package main

import (
	"context"
	"fmt"
)

func handlerLogin(state *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username required")
	}

	contextBackground := context.Background()
	name := cmd.args[0]

	_, err := state.db.GetUser(contextBackground, name)
	if err != nil {
		return fmt.Errorf("user %q is not registerd", name)
	}

	fmt.Println("Logging in...")

	err = state.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("%v logged in successfully!\n", name)

	return nil
}
