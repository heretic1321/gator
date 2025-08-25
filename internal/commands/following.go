package commands

import (
	"context"
	"fmt"
	"github.com/heretic1321/gator/internal/database"
)

func handleFollowing(state *State, args []string, user database.User)error {
	
	feeds, err := state.DB.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}

	if len(feeds) < 1 {
		fmt.Println("user is following no feed")
		return nil
	}
	fmt.Printf("%s is following the feeds below : \n", feeds[0].UserName)
	for _, feed := range feeds {
		fmt.Printf("%v\n", feed.FeedName)
	}
	
	return nil	
}
