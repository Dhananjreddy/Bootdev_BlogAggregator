package main

import (
	"log"
	"os"
	"github.com/Dhananjreddy/Bootdev_BlogAggregator/golang/internal/config"
)

type state struct {
	config *config.Config
}

func main()  {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	
	newState := &state{
		config: &cfg,
	}

	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	err = commands.run(newState, command{Name: commandName, Arguments: commandArgs})
	if err != nil {
		log.Fatal(err)
	}
}