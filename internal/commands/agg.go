package commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/heretic1321/gator/internal/database"
)

func scrapeFeeds(state *State) error {
	feed, err := state.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	if err := state.DB.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		return err
	}

	fmt.Printf("Fetching feed: %s (%s)\n", feed.Name, feed.Url)
	rss, err := state.Client.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rss.Channel.Item {
		publishedAt, _ := time.Parse(time.RFC1123Z, item.PubDate)
		if publishedAt.IsZero() {
			if t, err := time.Parse(time.RFC1123, item.PubDate); err == nil {
				publishedAt = t
			}
		}
		post := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: sql.NullString{String: item.Title, Valid: item.Title != ""},
			Url: item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: sql.NullTime{Time: publishedAt, Valid: !publishedAt.IsZero()},
			FeedID: feed.ID,
		}
		if err := state.DB.CreatePost(context.Background(), post); err != nil {
			fmt.Printf("save post error: %v\n", err)
		}
	}
	return nil
}

func handleAggregator(state *State, args []string) error{
	if len(args) < 1 {
		return fmt.Errorf("time_between_reqs argument required (e.g. 1s, 1m, 1h)")
	}
	timeBetweenRequests, err := time.ParseDuration(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		if err := scrapeFeeds(state); err != nil {
			fmt.Printf("scrape error: %v\n", err)
		}
	}
}
