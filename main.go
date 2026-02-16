package main

import (
	"fmt"
	"os"

	"github.com/kadlex-web/bloggator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	arguments []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	// if the length of arguments is 0 -- no user can be logged in
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("login expects a username")
	}
	username := cmd.arguments[0]
	err := s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("error setting username")
	}
	fmt.Printf("User has been set to %v\n", username)
	return nil
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
func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func main() {
	// create initial state
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config")
		os.Exit(1)
	}
	s := state{config: &cfg}
	// create initial commands struct
	commandsMap := commands{}
	// initialize map of possible commands
	commandsMap.commands = make(map[string]func(*state, command) error)
	//register login command
	commandsMap.register("login", handlerLogin)

	// grab the user input and check if enough arguments have been passed for a command
	input := os.Args
	if len(input) < 2 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}
	// if it's good, create a command struct and send to the commands struct to run
	userCmd := command{
		name: input[1],
		arguments: input[2:],
	}
	err = commandsMap.run(&s, userCmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
