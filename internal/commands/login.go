package commands

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(state *State, args []string ) error{
  if len(args) < 1 {
		return errors.New("a username field is required")
	}

	_, err := state.DB.GetUser(context.Background(), args[0])
	if err != nil {
		return err
	}

	state.Cfg.SetUser(args[0])
	fmt.Println("logged in successfully. Welcome ", args[0])
	return nil
}
