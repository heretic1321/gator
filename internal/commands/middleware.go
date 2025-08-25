package commands

import (
    "context"
    "errors"

    "github.com/heretic1321/gator/internal/database"
)

type loggedInHandler func(state *State, args []string, user database.User) error

func middlewareLoggedIn(next loggedInHandler) func(*State, []string) error {
    return func(s *State, args []string) error {
        if s == nil || s.Cfg == nil || s.DB == nil {
            return errors.New("application state not initialized")
        }
        if s.Cfg.CurrentUsername == "" {
            return errors.New("no user is logged in. run: login <username>")
        }

        user, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUsername)
        if err != nil {
            return err
        }

        return next(s, args, user)
    }
}


