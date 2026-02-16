package main

import (
	"fmt"
	"github.com/kadlex-web/bloggator/internal/config"
)

func main() {
	// read the config file
	cfg, _ := config.Read()
	fmt.Println(cfg)
	//update current user to "alex"
	cfg.SetUser("alex")
	//read the config file again
	cfg2, _ := config.Read()
	fmt.Println(cfg2)
}
