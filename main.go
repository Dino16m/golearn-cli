package main

import (
	"fmt"
	"os"
)

type Command string

const (
	Install Command = "install"
)

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("No arguments provided")
		os.Exit(1)
	}
	cmd := args[0]
	executeCommand(Command(cmd))
}

func executeCommand(cmd Command) {
	switch cmd {
	case Install:
		install()
		break
	default:
		fmt.Println("Unknown command ", cmd)
	}
}
