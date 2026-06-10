package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Vandush/Gator/internal/config"
	"github.com/Vandush/Gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	s := config.State{}
	c, _ := config.Read()
	s.Conf = &c

	c.DBURL = "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"

	db, err := sql.Open("postgres", c.DBURL)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	s.DB = database.New(db)

	cmds := config.Commands{}
	cmds.Make()
	cmds.Register("login", config.HandlerLogin)
	cmds.Register("register", config.HandlerRegister)
	cmds.Register("reset", config.HandlerResetUsers)
	cmds.Register("users", config.HandlerListUsers)

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("No command given.")
		os.Exit(1)
	}
	command := config.Command{
		Name:      args[0],
		Arguments: args[1:],
	}
	if err := cmds.Run(&s, command); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
