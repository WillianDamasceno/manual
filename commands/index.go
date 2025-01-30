package commands

import (
	"manual/helpers"
)

func Help() {
	helpers.Welcome()

	println("How to use:")
	println("  manual [command] [options]\n")

	println("Commands:")
	for _, command := range Commands {
		println("  ", command.Name, "-", command.Description)
	}
}

func Nix() {
	println("How to list all packages:")
	println("  nix-env -q")

	println()
	println("How uninstall a package:")
	println("  nix-env -e <package>")
}

func Apt() {
	println("How to list all packages")
	println("  apt list --installed")
}

func Git() {
	println("How set a config value globaly")
	println("  git config --global <config.name> <'value'>")
}
