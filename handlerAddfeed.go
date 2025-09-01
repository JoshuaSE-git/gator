package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JoshuaSE-git/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddfeed(state *State, cmd Command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("expected 2 arguments: feed name and url")
	}

	ctx := context.Background()
	id := uuid.New()
	currentTime := time.Now()
	feedName := cmd.args[0]
	feedUrl := cmd.args[1]

	params := database.CreateFeedParams{
		ID:        id,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      feedName,
		Url:       feedUrl,
	}

	feed, err := state.db.CreateFeed(ctx, params)
	if err != nil {
		return err
	}

	feedFollowId := uuid.New()
	userId := user.ID
	feedId := feed.ID

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        feedFollowId,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    userId,
		FeedID:    feedId,
	}

	_, err = state.db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feed)

	return nil
}
