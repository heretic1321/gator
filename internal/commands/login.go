package commands

import (
	"errors"
	"fmt"
)

func handlerLogin(state *State, args []string ) error{
  if len(args) < 1 {
		return errors.New("a username field is required")
	}
	state.Cfg.SetUser(args[0])
	fmt.Println("username set to :", args[0])
	return nil
}
