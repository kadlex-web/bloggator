package main

import "fmt"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

// method runs a given command with the provided state if it exists. method returns error value
func (c *commands) run(s *state, cmd command) error {
	// check the map to see if a command is registered
	val, ok := c.commands[cmd.name]
	if !ok {
		return fmt.Errorf("command does not exist. please register command and try again")
	}
	err := val(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

// method registers a new handler function for a command name. method has no return value
func (c *commands) registerCommand(name string, f func(*state, command) error) {
	c.commands[name] = f
}
