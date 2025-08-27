package main

import (
	"errors"
	"fmt"

	"github.com/JoshuaSE-git/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	cfg.SetUser("Joshua")

	cfg, err = config.Read()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", *cfg)
}

type Command struct {
	name string
	args []string
}

type Commands struct {
	commandMap map[string]func(state *State, cmd Command) error
}

func (c *Commands) register(name string, handler func(state *State, cmd Command) error) {
	c.commandMap[name] = handler
}

func handlerLogin(state *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return errors.New("too few arguments")
	}

	fmt.Println("Logging in...")

	user := cmd.args[0]
	err := state.Config.SetUser(user)
	if err != nil {
		return err
	}

	fmt.Printf("%v logged in successfully!\n", user)

	return nil
}
