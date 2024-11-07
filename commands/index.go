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

// TODO: implement self update
func Update() {
	helpers.NotImplemented("update")
}
