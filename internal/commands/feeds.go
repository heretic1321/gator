package commands
import (
	"context"
	"fmt"
)

 
func handleFeeds(state *State, args []string) error{
	feeds, err := state.DB.GetFeeds(context.Background())

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		
	  user, err := state.DB.GetUserById(context.Background(), feed.UserID)
		if err != nil{
			return err
		}
		fmt.Printf("Name: %s, URL: %s, User: %s\n", feed.Name, feed.Url, user.Name)
	}

	return nil
}	
