package main

import (
	"context"
	"fmt"

	"github.com/kadlex-web/bloggator/internal/database"
)

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("unfollow command only takes 1 additional argument. ex: unfollow http://rssfeed.io")
	}
	feedURL := cmd.arguments[0]
	userID := user.ID
	feedID, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed ID")
	}
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: userID, FeedID: feedID})
	if err != nil {
		return fmt.Errorf("error unfollowing feed")
	}
	return nil
}
