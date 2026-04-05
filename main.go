package main

import (
	"fmt"
//	"os"
	"github.com/Vandush/Gator/internal/config"
)

func main() {
	conf, _ := config.Read()
	conf.SetUser("Vandush")
	fmt.Printf("%v", conf)
//	data, _ := os.ReadFile(config.ConfigPath())
//	fmt.Println(string(data))
}
