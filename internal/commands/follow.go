package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/heretic1321/gator/internal/database"
)

func handleFollow(state *State, args []string) error {
	if len(args) < 1{
		return errors.New("url field missing from arguments")
	}

	url := args[0]

	user, err := state.DB.GetUser(context.Background(), state.Cfg.CurrentUsername)
	if err != nil {
		return err
	}

	feed, err := state.DB.GetFeedByUrl(context.Background(), url)
	

	if err != nil {
		err = handleAddfeed(state, []string{url})
		if err != nil {
			return err
		}

		newFeed, err := state.DB.GetFeedByUrl(context.Background(), url) 
		if err != nil {
			return err
		}

		feed = newFeed
		
	}

	feedFollowOpts := database.CreateFeedFollowParams{
		ID: uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
	}


	_, err = state.DB.CreateFeedFollow(context.Background(), feedFollowOpts)	
	fmt.Printf("User %s, started following feed: %s", user.Name, feed.Name)

	return nil
}	
