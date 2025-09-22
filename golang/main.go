package main

import (
	"fmt"
	"log"
	"github.com/Dhananjreddy/Bootdev_BlogAggregator/golang/internal/config"
)

type state struct {
	cfg *Config
}

func main()  {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	err = cfg.SetUser("Dhananjay")
	if err != nil {
		log.Fatalf("error setting username in config: %v", err)
	}
	fmt.Println("DBURL: %s", cfg.DBURL)
	fmt.Println("CurrentUserName: %s", cfg.CurrentUserName)
}