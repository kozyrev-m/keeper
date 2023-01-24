package main

import (
	"log"
	"github.com/kozyrev-m/keeper/internal/agent"
)

func main() {
	if err := agent.StartClient(); err != nil {
		log.Fatal(err)
	}
}