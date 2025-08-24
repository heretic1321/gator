package cli

import (
	"errors"

	"github.com/heretic1321/gator/internal/commands"
)

type App struct {
	state *commands.State	
}

func (a *App) Run( args []string) error{	
	if len(args) < 1 {
		return errors.New("not enough arguments were provided")
	}
	cmds := commands.New()
	
	err := cmds.Run(a.state, args)
	if err != nil {
		return err
	}
	return nil
}

func New (state commands.State) App{	
	
	return App{
		state: &state,
	}
}
