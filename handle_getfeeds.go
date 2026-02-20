package main

import (
	"context"
	"fmt"
)

func handlerGetFeeds(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("feeds command does not take any arguments")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	// for each user, display its name in the console
	if len(feeds) == 0 {
		fmt.Println("No users registered. Please use 'register <name>' to register a user.")
	}
	for i, feed := range feeds {
		feedCreator, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("can't find creator of feed")
		}
		fmt.Printf("Feed %v\n\n", i)
		fmt.Printf("Name of Feed: %v\n", feed.Name)
		fmt.Printf("URL of Feed: %v\n", feed.Url)
		fmt.Printf("Created by: %v\n", feedCreator)
		fmt.Println("----------------------------------------")
	}
	return nil
}
