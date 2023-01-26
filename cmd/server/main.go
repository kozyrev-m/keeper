package main

import (
	"log"

	"github.com/kozyrev-m/keeper/internal/master/config"
	"github.com/kozyrev-m/keeper/internal/master/server"
)

func main() {
	log.Println("Hello!")

	if err := server.StartServer(config.NewConfig()); err != nil {
		log.Fatal(err)
	}
}