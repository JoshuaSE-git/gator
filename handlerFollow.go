package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/JoshuaSE-git/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(state *State, cmd Command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("url is required")
	}

	ctx := context.Background()
	url := cmd.args[0]

	feed, err := state.db.GetFeed(ctx, url)
	if err != nil {
		return err
	}

	id := uuid.New()
	currentTime := time.Now()
	userId := user.ID
	feedId := feed.ID

	params := database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    userId,
		FeedID:    feedId,
	}

	feedFollow, err := state.db.CreateFeedFollow(ctx, params)
	if err != nil {
		return err
	}

	fmt.Printf("%s followed %q\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
