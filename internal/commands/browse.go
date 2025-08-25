package commands

import (
    "context"
    "fmt"
    "strconv"

    "github.com/heretic1321/gator/internal/database"
)

func handleBrowse(state *State, args []string, user database.User) error {
    limit := int32(2)
    if len(args) >= 1 {
        if v, err := strconv.Atoi(args[0]); err == nil && v > 0 {
            limit = int32(v)
        }
    }
		

    posts, err := state.DB.GetPostsForUser(context.Background(),database.GetPostsForUserParams{ID: user.ID, Limit: limit})
    if err != nil {
        return err
    }

    if len(posts) == 0 {
        fmt.Println("No posts found.")
        return nil
    }

    for _, p := range posts {
        title := p.Title.String
        if title == "" {
            title = "(no title)"
        }
        fmt.Printf("- %s\n  %s\n", title, p.Url)
    }
    return nil
}


