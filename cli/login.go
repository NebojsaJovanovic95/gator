package cli

import (
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("username required")
	}

	username := cmd.Args[0]

	err := s.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User set to %s\n", username)
	return nil
}
