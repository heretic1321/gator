package commands

import "context"


func handleReset(state *State, args []string) error{
	err := state.DB.DeleteUsers(context.Background())

	if err != nil {
		return err
	}

	return nil
}
