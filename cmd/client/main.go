package main

import (
	"log"

	"github.com/kozyrev-m/keeper/internal/agent/client"
)

func main() {
	if err := client.StartClient(); err != nil {
		log.Fatal(err)
	}
}