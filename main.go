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

	switch os.Args[1] {
	case "help", "-h", "--help":
		commands.Help()
	case "update":
		commands.Update()
	case "nix", "linux":
		helpers.NotImplemented(os.Args[1])
	default:
		helpers.NotFound(os.Args[1])
	}
}
