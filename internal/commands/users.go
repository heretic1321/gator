package commands

import (
	"context"
	"fmt"
)


func handleUsers(state *State, args []string) error {
	
	users, err := state.DB.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _,user := range users{
		printableName := "* " + user.Name
		if user.Name == state.Cfg.CurrentUsername {
			printableName += " (current)"
		}
		fmt.Println(printableName)	
	}
	return nil
}
