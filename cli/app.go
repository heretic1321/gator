package cli

import (
	"errors"
	"strings"

	"github.com/heretic1321/gator/internal/commands"
	"github.com/heretic1321/gator/internal/config"
)

type App struct {
	Conf *config.Config
}


func (a *App) Run( args []string) error{
	err := a.Conf.SetUser("heretic")	
	if err != nil {
		return err
	}
	if len(args) < 1 {
		return errors.New("not enough arguments were provided")
	}
	cmds := commands.New()
	
	command := commands.Command{
		Name: strings.ToLower(args[0]),
		Args: args[1:],
	}
	err = cmds.Run(a.Conf, command)
	if err != nil {
		return err
	}
	return nil
}

func New (conf *config.Config) App{
	return App{
		Conf : conf,
	}
}
