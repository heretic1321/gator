package commands

import (
    "context"
    "errors"
    "fmt"

    "github.com/heretic1321/gator/internal/database"
)

// handleUnfollow removes a feed follow by URL for the currently logged-in user.
func handleUnfollow(state *State, args []string, user database.User) error {
    if len(args) < 1 {
        return errors.New("url field missing from arguments")
    }
    url := args[0]

    // Ensure feed exists (optional). We can rely on delete to no-op if missing,
    // but fetching helps produce more informative output.
    _, err := state.DB.GetFeedByUrl(context.Background(), url)
    if err != nil {
        return fmt.Errorf("feed with url %s not found", url)
    }

    if err := state.DB.DeleteFeedFollowByUserAndURL(context.Background(), database.DeleteFeedFollowByUserAndURLParams{
			UserID: user.ID,
			Url: url,
		}); err != nil {
        return err
    }

    fmt.Printf("User %s unfollowed feed %s\n", user.Name, url)
    return nil
}


