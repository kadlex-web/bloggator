package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/kadlex-web/bloggator/internal/config"
	"github.com/kadlex-web/bloggator/internal/database"
	_ "github.com/lib/pq" // side effects import
)

type state struct {
	db *database.Queries
	config *config.Config
}

func main() {

	// create initial state
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config")
		os.Exit(1)
	}
	s := state{config: &cfg}
	// open a connection to the database using your dbURL which was stored in the state config
	db, err := sql.Open("postgres", s.config.Dburl)
	// save the database within the program state
	s.db = database.New(db)

	// create initial commands struct
	commandsMap := commands{}
	// initialize map of possible commands
	commandsMap.commands = make(map[string]func(*state, command) error)
	//register login command
	commandsMap.registerCommand("login", handlerLogin)
	// register the register command
	commandsMap.registerCommand("register", handlerRegister)

	// grab the user input and check if enough arguments have been passed for a command
	input := os.Args
	if len(input) < 2 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}
	// if it's good, create a command struct and send to the commands struct to run
	userCmd := command{
		name:      input[1],
		arguments: input[2:],
	}
	err = commandsMap.run(&s, userCmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
