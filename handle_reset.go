package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("reset doesn't take arguments")
	}
	err := s.db.Reset(context.Background())
	return err
}
