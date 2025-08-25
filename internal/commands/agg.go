package commands

import (
	"context"
	"fmt"
)

func handleAggregator(state *State, args []string) error{
	feed, err := state.Client.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
	return err
	}

	fmt.Printf("%+v\n", *feed)
	return nil
}
