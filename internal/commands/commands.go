package commands

import (
	"errors"

	"github.com/heretic1321/gator/internal/config"
)

type Command struct {
	Name string
	Args []string
}


type CommandCallbackRegistry struct {
	Reg map[string]func(*config.Config, Command) error
}

func New() CommandCallbackRegistry {
	cmds := CommandCallbackRegistry{Reg: make(map[string]func(*config.Config, Command)error )}
	cmds.Register("login", handlerLogin)
	
	return cmds
}

func (c *CommandCallbackRegistry) Run(conf *config.Config, cmd Command) error{	

	callback,ok := c.Reg[cmd.Name]
	if !ok {
		return errors.New("invalid command. run help to show the commands")
	}

	err := callback(conf , cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *CommandCallbackRegistry) Register(name string, f func(*config.Config,Command)error) {
	
	if c.Reg == nil {
		c.Reg = make(map[string]func(*config.Config, Command) error)
	}

	c.Reg[name] = f
}



