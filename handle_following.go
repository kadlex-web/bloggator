package main

import (
	"context"
	"fmt"

	"github.com/kadlex-web/bloggator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("following command takes no arguments")
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching follows from the database")
	}
	if len(feedFollows) == 0 {
		fmt.Println("You haven't followed any feeds yet. Please use 'follow <url> to follow a feed")
		return nil
	}
	fmt.Printf("User: %s is following these feeds:\n", user.Name)
	for _, feed := range feedFollows {
		fmt.Printf("- %s\n", feed.FeedName)
	}
	return nil
}
