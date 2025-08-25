package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/heretic1321/gator/internal/database"
)

func handleAddfeed(state *State, args []string) error{
	if len(args) < 2{
		return errors.New("2 required arguments missing [name, url]")
	} else if len(args)< 1 {
		return errors.New("1 required arguments missing [name, url]")
	}
	name := args[0]
	url := args[1]
	user, err := state.DB.GetUser(context.Background(), state.Cfg.CurrentUsername)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	createFeedParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user.ID,
	}
	feed, err := state.DB.CreateFeed(context.Background(),createFeedParams)

	if err != nil{
		return err
	}

	err = handleFollow(state, []string{url})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow relation, %s\n", err.Error())
	} 

	fmt.Printf("%v\n", feed)
	return nil
}
