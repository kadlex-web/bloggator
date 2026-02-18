package main

import (
	"context"
	"fmt"
)

// uses sqlc generated go code to get all the users from the table and print them to the terminal
func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("users command does not take any arguments")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	// for each user, display its name in the console
	if len(users) == 0 {
		fmt.Println("No users registered. Please use 'register <name>' to register a user.")
	}
	// get the value of the current logged in user and add (current) after their name in the user list
	currentUser := s.config.Username
	for _, user := range users {
		if currentUser == user.Name {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}
	return nil
}
