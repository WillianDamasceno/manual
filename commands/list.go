package commands

type Command struct {
	Name        string
	Description string
	Run         func()
	Registered  bool
}

var Commands = map[string]Command{
	"update": {
		Name:        "update",
		Description: "Updates manual to the latest version",
		Run:         Update,
		Registered:  true,
	},
	"nix": {
		Name:        "nix",
		Description: "Nix common tasks",
		Run:         Nix,
		Registered:  true,
	},
}
