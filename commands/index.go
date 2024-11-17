package commands

import "manual/helpers"

func Help() {
	helpers.Welcome()

	println("How to use:")
	println("  manual [command] [options]\n")

	println("Commands:")
	for _, command := range Commands {
		println("  ", command.Name, "-", command.Description)
	}
}

func Update() {
	helpers.NotImplemented("update")
}

func Nix() {
	helpers.NotImplemented("update")
}
