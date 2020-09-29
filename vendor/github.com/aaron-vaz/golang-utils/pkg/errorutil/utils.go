package errorutil

import (
	"log"
	"os"
	"runtime/debug"
)

// ErrCheck is a helper method to check is a error object is a valid error
// if the error is a valid it will log it
func ErrCheck(err error, exit bool) {
	if err != nil {
		msg := err.Error()

		if stack := os.Getenv("DUMP_STACK"); stack == "true" {
			msg += "\n" + string(debug.Stack())
		}

		log.Println(msg)

		if exit {
			os.Exit(1)
		}
	}
}
