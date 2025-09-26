package main

import(
	"fmt"
	"time"
	"context"
	"github.com/Dhananjreddy/Bootdev_BlogAggregator/golang/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Arguments[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error fetching feed from database: %w", err)
	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error fetching user from database: %w", err)
	}

	followRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("Error Creating feed follow in database: %w", err)
	}
	printFeedFollow(followRow.UserName, followRow.FeedName)
	return nil
}

func handlerListFeedFollows(s *state, cmd command) error {
	name := s.config.CurrentUserName
	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("Error Fetching user from database: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error Fetching Feed follows for user from database: %w", err)
	}

	for _, feed := range feedFollows{
		fmt.Printf("* %s\n", feed.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	feedUrl := cmd.Arguments[0]

	name := s.config.CurrentUserName
	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("Error Fetching user from database: %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("Error Fetching feed from database: %w", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Error Unfollowing feed from database: %w", err)
	}

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}