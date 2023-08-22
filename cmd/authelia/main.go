package main

import (
	"os"

	"github.com/authelia/authelia/v4/pkg/commands"
)

func main() {
	if err := commands.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
