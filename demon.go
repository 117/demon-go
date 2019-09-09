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
}

func main() {
	if len(os.Args) <= 1 {
		command.GetCommand().Help()
		return
	}
	command.GetCommand().Execute()
}
