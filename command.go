package main

import "fmt"

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

func (c *Commands) run(state *State, cmd Command) error {
	handler, ok := c.commandMap[cmd.name]
	if !ok {
		return fmt.Errorf("%q command doesn't exist", cmd.name)
	}

	err := handler(state, cmd)
	if err != nil {
		return err
	}

	return nil
}
