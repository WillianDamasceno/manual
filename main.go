package main

import (
	"manual/commands"
	"manual/helpers"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		helpers.Welcome()
		helpers.WarnAboutHelp()
		return
	}

	command, exists := commands.Commands[os.Args[1]]

	// Call help
	switch os.Args[1] {
	case "help", "-h", "--help":
		commands.Help()
		return
	}

	// When command is not found
	if !exists {
		helpers.NotFound(os.Args[1])
		return
	}

	// When command is not registered
	if !command.Registered {
		helpers.NotImplemented(os.Args[1])
		return
	}

	helpers.Welcome()
	command.Run()
}
