package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kadlex-web/bloggator/internal/database"
)

// HandlerFollow implements the follow command
func handlerFollow(s *state, cmd command, user database.User) error {
	// if the user passed more then one argument to the function -- print the error below
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("follow command only requires url. ex: follow https://hnrss.org/newest")
	}
	// URL passed to the function and used to run query for feedID
	feedURL := cmd.arguments[0]
	feedID, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed id from database")
	}
	u := uuid.New()
	feedParams := database.CreateFeedFollowParams{
		ID:        u,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID_2:      user.ID,
		ID_3:      feedID,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error creating feed follow record in feeds table")
	}
	fmt.Printf("User: %s is now following Feed: %s\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}
