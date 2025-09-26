package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Deleted users: %w")
	return nil
}