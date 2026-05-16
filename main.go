package main

import (
	"fmt"
	"os"
	"database/sql"
	"github.com/Vandush/Gator/internal/config"
	"github.com/Vandush/Gator/internal/database"
)

import _ "github.com/lib/pq"

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

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	s.Db = database.New(db)

	os.Exit(0)
	//	conf.SetUser("Vandush")
}
