package helpers

import (
	"fmt"
	"manual/src"
	"os"
	"path/filepath"
)

func Welcome() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Manual", src.Version, filepath.Dir(exePath), "\n")
}

func WarnAboutHelp() {
	println("Run 'manual help' for more information.")
}

func NotFound(command string) {
	fmt.Println("Command '", command, "' not found.")
	WarnAboutHelp()
	os.Exit(1)
}

func NotImplemented(command string) {
	println("Sorry, command '", command, "' is not implemented yet.")
	println("Run 'manual update' to download to the latest version.")
	os.Exit(1)
}
