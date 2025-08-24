package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/heretic1321/gator/internal/database"
)

func handleRegister(state *State, args []string) error{
	
	if len(args) < 1  {
		return errors.New("name not specified for user")
	}

	user := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: args[0],
	}

	_, err := state.DB.CreateUser(context.Background(), user)
	
	if err != nil {
		return err;
	}

	err = state.Cfg.SetUser(args[0])
	if err != nil {
		return err
	}	

	fmt.Println("user registered successfully")

	return nil
}
