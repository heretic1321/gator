package commands

import (
	"context"
	"errors"
	"fmt"
)

func handleFollowing(state *State, args []string)error {
	
	user, err := state.DB.GetUser(context.Background(), state.Cfg.CurrentUsername)

	if err != nil {
		return err
	}

	feeds, err := state.DB.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return errors.New("user is not following any feeds")
	}

	if len(feeds) < 1 {
		return errors.New("user is following no feed")
	}
	fmt.Printf("%s is following the feeds below : \n", feeds[0].UserName)
	for _, feed := range feeds {
		fmt.Printf("%v\n", feed.FeedName)
	}
	
	return nil	
}
