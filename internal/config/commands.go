package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Vandush/Gator/internal/database"
	"github.com/google/uuid"
)

type State struct {
	DB   *database.Queries
	Conf *Config
}

type Command struct {
	Name      string
	Arguments []string
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("no input was given")
	}
	if len(cmd.Arguments) > 1 {
		return fmt.Errorf("too many arguments")
	}

	ctx := context.Background()
	if _, err := s.DB.GetUser(ctx, cmd.Arguments[0]); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		} else {
			return fmt.Errorf("%v is not a registered user", cmd.Arguments[0])
		}
	}

	if err := s.Conf.SetUser(cmd.Arguments[0]); err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Printf("%s has logged in successfully.\n", cmd.Arguments[0])
	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("no input was given")
	}
	if len(cmd.Arguments) > 1 {
		return fmt.Errorf("too many arguments")
	}

	ctx := context.Background()
	arg := database.CreateUserParams{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Arguments[0],
	}

	// This should return a "no rows" error, else there is actual error or user is found.
	if _, err := s.DB.GetUser(ctx, arg.Name); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	} else {
		return fmt.Errorf("%v is already registered", arg.Name)
	}

	if _, err := s.DB.CreateUser(ctx, arg); err != nil {
		return err
	}

	if err := s.Conf.SetUser(arg.Name); err != nil {
		return fmt.Errorf("%v", err)
	}

	fmt.Printf("%v has successfully been registered.\n", arg.Name)
	return nil
}

func HandlerResetUsers(s *State, cmd Command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("'reset' does not take arguments")
	}

	ctx := context.Background()

	if err := s.DB.DropUserTable(ctx); err != nil {
		return err
	}

	if err := s.DB.CreateUserTable(ctx); err != nil {
		return err
	}

	return nil
}

func HandlerListUsers(s *State, cmd Command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("'users' does not take arguments")
	}

	ctx := context.Background()
	users, err := s.DB.GetUsers(ctx)
	if err != nil {
		return err
	}

	for i := range len(users) {
		if users[i] == s.Conf.CurrentUserName {
			fmt.Printf("%v (current)\n", users[i])
		} else {
			fmt.Printf("%v\n", users[i])
		}
	}
	return nil
}

type Commands struct {
	commands map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	if r, exists := c.commands[cmd.Name]; exists {
		if err := r(s, cmd); err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("'%s' is not a command", cmd.Name)
	}
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.commands[name] = f
}

func (c *Commands) Make() {
	c.commands = make(map[string]func(*State, Command) error)
}
