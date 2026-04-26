package config

import (
	"fmt"
)

type State struct {
	Conf *Config
}

type Command struct {
	Name string
	Arguments []string
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("No input was given.\n")
	}
	if len(cmd.Arguments) > 1 {
		return fmt.Errorf("Too many arguments.\n")
	}
	if err := s.Conf.SetUser(cmd.Arguments[0]); err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Printf("%s has logged in successfully.\n", cmd.Arguments[0])
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
	return
}

func (c *Commands) Make() {
	c.commands = make(map[string]func(*State, Command) error)
	return
}
