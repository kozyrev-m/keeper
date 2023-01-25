package main

import (
	"log"
	"github.com/kozyrev-m/keeper/internal/agent/terminal"
)

func main() {
	if err := terminal.NewTerminal().StartTerminal(); err != nil {
		log.Fatal(err)
	}
}