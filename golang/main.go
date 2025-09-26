package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"context"
	"github.com/Dhananjreddy/Bootdev_BlogAggregator/golang/internal/config"
	"github.com/Dhananjreddy/Bootdev_BlogAggregator/golang/internal/database"
)

type state struct {
	db  *database.Queries
	config *config.Config
}

func main()  {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	newState := &state{
		db: dbQueries,
		config: &cfg,
	}

	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerListUsers)
	commands.register("agg", handlerAggregator)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerListFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerListFeedFollows))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.register("browse", middlewareLoggedIn(handlerBrowse))

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

func middlewareLoggedIn(handler func(s *state, cmd command) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		_, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd)
	}
}