package commands

type Command struct {
	Name        string
	Description string
	Run         func()
}

var Commands = map[string]Command{
	"update": {
		Name:        "update",
		Description: "Updates manual to the latest version",
		Run:         Update,
	},
}
