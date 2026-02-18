package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kadlex-web/bloggator/internal/database"
)

// updates the user in the config file (login for the flat file)
func handlerLogin(s *state, cmd command) error {
	// if the length of arguments isn't 1, invalid command has been passed
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("cli command is: login <name>")
	}
	// check if user exists in db, if not, user is not logged in
	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("user does not exist in the db. please register before trying again")
	}
	username := cmd.arguments[0]
	err = s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("error setting username")
	}
	fmt.Printf("User has been set to %v\n", username)
	return nil
}

// registers users to the db; so they may login and use features
func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf(("improper use of register. proper -- register <name>"))
	}
	u := uuid.New()
	// if GetUser does not return an error, then the user already exists in the db and therefore cannot be created
	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err == nil {
		return fmt.Errorf("user already exists. all users must be unique.")
	}
	// if getUser returns an err -- it means the user is okay to be created!
	newUser := database.CreateUserParams{
		ID:        u,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
	}
	// create a new user in the table
	data, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("issue creating user in database. check to make sure database is active")
	}
	// update the current user in the config
	s.config.SetUser(cmd.arguments[0])
	fmt.Println("User was successfully created")
	fmt.Println(data)
	return nil
}
