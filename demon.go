package main

import (
	"os"

	"./command"
)

func main() {
	if len(os.Args) <= 1 {
		command.GetBase().Help()
		return
	}
	command.GetBase().Execute()
}
