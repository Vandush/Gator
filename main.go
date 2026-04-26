package main

import (
	"fmt"
	"os"
	"github.com/Vandush/Gator/internal/config"
)

func main() {
	s := config.State{}
	c, _ := config.Read()
	s.Conf = &c
	cmds := config.Commands{}
	cmds.Make()
	cmds.Register("login", config.HandlerLogin)
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("No command given.")
		os.Exit(1)
	}
	command := config.Command{
		Name: args[0],
		Arguments: args[1:],
	}
	if err := cmds.Run(&s, command); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
	//	conf.SetUser("Vandush")
}
