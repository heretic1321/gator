package commands

import (
	"errors"
)

type Command struct {
	Name string
	Args []string
}


type CommandCallbackRegistry struct {
	Reg map[string]func(*State, []string) error
}

func New() CommandCallbackRegistry {
	cmds := CommandCallbackRegistry{Reg: make(map[string]func(*State, []string)error )}
	cmds.Register("login", handlerLogin)
	cmds.Register("register", handleRegister)
	cmds.Register("reset", handleReset)
	cmds.Register("users", handleUsers)
	return cmds
}

func (c *CommandCallbackRegistry) Run(state *State, args []string) error{	

	callback,ok := c.Reg[args[0]]
	if !ok {
		return errors.New("invalid command. run help to show the commands")
	}
	// since first arg is the command name 
	err := callback(state, args[1:])
	if err != nil {
		return err
	}

	return nil
}

func (c *CommandCallbackRegistry) Register(name string, f func(*State,[]string)error) {
	
	if c.Reg == nil {
		c.Reg = make(map[string]func(*State, []string) error)
	}

	c.Reg[name] = f
}



