package cli

import "github.com/NebojsaJovanovic95/gator.git/internal/config"

type State struct {
	cfg *config.Config
}

func (s *State) Config() *config.Config {
	return s.cfg
}

func (s *State) SetUser(username string) error {
	return s.cfg.SetUser(username)
}

func NewState(cfg *config.Config) *State {
	return &State{cfg: cfg}
}
