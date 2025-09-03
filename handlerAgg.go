package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/JoshuaSE-git/gator/internal/database"
	"github.com/JoshuaSE-git/gator/internal/rss"
	"github.com/google/uuid"
)

const URL = "https://www.wagslane.dev/index.xml"

func handlerAgg(state *State, cmd Command) error {
	if len(cmd.args) < 1 {
		return errors.New("duration string required")
	}

	ctx := context.Background()
	interval, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(interval)

	fmt.Printf("Collecting feeds every %s\n", interval.String())
	for ; ; <-ticker.C {
		err := scrapeFeeds(state, ctx)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(state *State, ctx context.Context) error {
	feeds, err := state.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		feedId := feed.ID
		currentTime := sql.NullTime{Time: time.Now(), Valid: true}
		params := database.MarkFeedFetchedParams{
			ID:            feedId,
			LastFetchedAt: currentTime,
		}

		err = state.db.MarkFeedFetched(ctx, params)
		if err != nil {
			return err
		}

		feedUrl := feed.Url

		rssFeed, err := rss.FetchFeed(ctx, feedUrl)
		if err != nil {
			return err
		}

		for _, item := range rssFeed.Channel.Item {
			postId := uuid.New()
			currentTime := time.Now()
			publishDate, err := time.Parse(time.RFC1123Z, item.PubDate)
			if err != nil {
				return err
			}
			postParams := database.CreatePostParams{
				ID:          postId,
				CreatedAt:   currentTime,
				UpdatedAt:   currentTime,
				Title:       item.Title,
				Url:         item.Link,
				Description: item.Description,
				PublishedAt: publishDate,
				FeedID:      feedId,
			}

			err = state.db.CreatePost(ctx, postParams)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
