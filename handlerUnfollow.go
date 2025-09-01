package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/JoshuaSE-git/gator/internal/database"
)

func handlerUnfollow(state *State, cmd Command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("url is required for unfollow")
	}

	ctx := context.Background()
	feedUrl := cmd.args[0]

	feed, err := state.db.GetFeed(ctx, feedUrl)
	if err != nil {
		return err
	}

	userId := user.ID
	feedId := feed.ID
	params := database.DeleteFeedFollowParams{
		UserID: userId,
		FeedID: feedId,
	}

	err = state.db.DeleteFeedFollow(ctx, params)
	if err != nil {
		return err
	}

	fmt.Printf("unfollowed %q\n", feed.Name)

	return nil
}
