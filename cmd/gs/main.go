package main

import (
	"fmt"
	"os"

	"github.com/go-gulfstream/gs/internal/commands"
)

const version = "v0.0.0"

func main() {
	app, err := commands.New(version)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	if err := app.Execute(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
