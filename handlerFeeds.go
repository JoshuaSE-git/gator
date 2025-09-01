package main

import (
	"context"
	"fmt"
)

func handlerFeeds(state *State, cmd Command) error {
	ctx := context.Background()

	feeds, err := state.db.GetFeeds(ctx)
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("no feeds available...")
	}

	for _, feed := range feeds {
		fmt.Printf(" - %q: %s\n", feed.Name, feed.Url)
	}

	return nil
}
