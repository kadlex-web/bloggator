package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("following command takes no arguments")
	}
	currentUser, err := s.db.GetUser(context.Background(), s.config.Username)
	if err != nil {
		return fmt.Errorf("error fetching current user from config")
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("error fetching follows from the database")
	}
	if len(feedFollows) == 0 {
		fmt.Println("You haven't followed any feeds yet. Please use 'follow <url> to follow a feed")
		return nil
	}
	fmt.Printf("User: %s is following these feeds:\n", currentUser.Name)
	for _, feed := range feedFollows {
		fmt.Printf("- %s\n", feed.FeedName)
	}
	return nil
}
