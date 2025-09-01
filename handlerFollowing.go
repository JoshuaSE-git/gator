package main

import (
	"context"
	"fmt"

	"github.com/JoshuaSE-git/gator/internal/database"
)

func handlerFollowing(state *State, cmd Command, user database.User) error {
	userName := user.Name
	ctx := context.Background()

	feedFollows, err := state.db.GetFeedFollowsUser(ctx, userName)
	if err != nil {
		return err
	}

	if len(feedFollows) == 0 {
		fmt.Println("Not following any feeds...")
	}

	for _, feedFollow := range feedFollows {
		fmt.Printf(" - %q: %s\n", feedFollow.FeedName, feedFollow.FeedUrl)
	}

	return nil
}
