package main

import (
	"fmt"
	"os"

	"github.com/117/logger"

	"./command"
)

func init() {
	logger.Formatter(func(level logger.Level, message string, vars ...interface{}) string {
		return fmt.Sprintf("demon > %s", fmt.Sprintf(message, vars...))
	})

	// dir, err := os.UserCacheDir()
	// make
}

func main() {
	if len(os.Args) <= 1 {
		command.GetCommand().Help()
	} else {
		command.GetCommand().Execute()
	}
}
