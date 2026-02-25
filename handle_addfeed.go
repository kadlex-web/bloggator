package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kadlex-web/bloggator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("addfeed command takes two arguments. ex: addfeed google www.google.com")
	}
	// get the currently logged-in user from the database
	currentUser, err := s.db.GetUser(context.Background(), s.config.Username)
	if err != nil {
		return fmt.Errorf("cannot get currently logged in user")
	}
	feedName := cmd.arguments[0]
	feedUrl := cmd.arguments[1]
	u := uuid.New()
	feedParams := database.CreateFeedParams{
		ID:        u,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    currentUser.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}
	fmt.Println(feed.ID)
	fmt.Println(feed.CreatedAt)
	fmt.Println(feed.UpdatedAt)
	fmt.Println(feed.Name)
	fmt.Println(feed.Url)
	fmt.Println(feed.UserID)

	followParams := database.CreateFeedFollowParams{
		ID:        u,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID_2:      currentUser.ID,
		ID_3:      u,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("error creating feed follow record in feeds table")
	}
	fmt.Printf("User: %s is now following Feed: %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
