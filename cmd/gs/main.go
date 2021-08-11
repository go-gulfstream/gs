package main

import (
	"fmt"
	"os"

	"github.com/go-gulfstream/gs/internal/commands"
)

var buildVersion = "v0.0.0"

func main() {
	app, err := commands.New(buildVersion)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	if err := app.Execute(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
