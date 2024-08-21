package main

import (
	"log"
	"os"

	"github.com/ohohestudio/texcoco/commands"
)

func main() {
	log.SetFlags(0)

	err := commands.Execute(os.Args[1:])
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
