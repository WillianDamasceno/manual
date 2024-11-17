package helpers

import (
	"fmt"
	"os"
)

func Error(err error, messages ...interface{}) {
	if err != nil {
		// Append the error to the messages
		messages = append(messages, err)

		// Print all the messages
		fmt.Println(messages...)
		os.Exit(1)
	}
}

func PanicIf(value bool, messages ...interface{}) {
	if value {
		panic(messages)
	}
}
