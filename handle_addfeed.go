package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kadlex-web/bloggator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("addfeed command takes two arguments. ex: addfeed google www.google.com")
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
		UserID:    user.ID,
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
		ID_2:      user.ID,
		ID_3:      u,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("error creating feed follow record in feeds table")
	}
	fmt.Printf("User: %s is now following Feed: %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
