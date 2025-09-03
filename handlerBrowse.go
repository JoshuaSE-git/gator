package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JoshuaSE-git/gator/internal/database"
)

func handlerBrowse(state *State, cmd Command, user database.User) error {
	limit := int32(2)
	if len(cmd.args) > 0 {
		num, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = int32(num)
	}

	ctx := context.Background()
	params := database.GetPostsParams{
		Name:  user.Name,
		Limit: limit,
	}

	posts, err := state.db.GetPosts(ctx, params)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("Title: %q (%s)\n", post.Title, post.FeedName)
		fmt.Printf("Published: %s\n", post.PublishedAt.String())
		fmt.Printf("Description: %q\n", post.Description)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Println()
	}

	return nil
}
