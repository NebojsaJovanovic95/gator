package main

import (
	"fmt"
	"os"

	"github.com/NebojsaJovanovic95/gator.git/cli"
	"github.com/NebojsaJovanovic95/gator.git/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	state := cli.NewState(&cfg)

	commands := cli.Commands{
		Handlers: make(map[string]func(*cli.State, cli.Command) error),
	}
	commands.Register("login", cli.HandlerLogin)

	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := commands.Run(state, cmd); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
