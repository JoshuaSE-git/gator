package main

import (
	"context"
	"fmt"

	"github.com/JoshuaSE-git/gator/internal/rss"
)

const URL = "https://www.wagslane.dev/index.xml"

func handlerAgg(state *State, cmd Command) error {
	ctx := context.Background()

	rssFeed, err := rss.FetchFeed(ctx, URL)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", *rssFeed)

	return nil
}
