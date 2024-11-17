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
		println("Error:", err)
	}

	println("Manual", src.Version, filepath.Dir(exePath))
	println()
}

func WarnAboutHelp() {
	println("Run 'manual help' for more information.")
}

func NotFound(command string) {
	print("Command '")
	print(command)
	println("' not found.")
	WarnAboutHelp()
	os.Exit(1)
}

func NotImplemented(command string) {
	print("Sorry, command '")
	print(command)
	println("' is not implemented yet.")
	println("Run 'manual update' to download to the latest version.")
	os.Exit(1)
}

func CheckSudo() {
	// Check if the program is running as root
	if os.Geteuid() != 0 {
		fmt.Println("This operation requires elevated permissions. Please run the program with `sudo`.")
		os.Exit(1)
	}
}
