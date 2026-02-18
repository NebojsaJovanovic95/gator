package main

import (
	"context"
	"fmt"

	"github.com/NebojsaJovanovic95/gator.git/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {

		if s.cfg.CurrentUserName == "" {
			return fmt.Errorf("user not selected")
		}

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't get user %s: %w", s.cfg.CurrentUserName, err)
		}

		return handler(s, cmd, user)
	}
}
