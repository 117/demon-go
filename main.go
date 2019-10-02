package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/117/logger"
	"github.com/streamwithme/demon/commands"
)

const (
	usage = ` _____   ______  __    __  ______  __   __    
/\  __-./\  ___\/\ "-./  \/\  __ \/\ "-.\ \   
\ \ \/\ \ \  __\\ \ \-./\ \ \ \/\ \ \ \-.  \  
 \ \____-\ \_____\ \_\ \ \_\ \_____\ \_\\"\_\ 
  \/____/ \/_____/\/_/  \/_/\/_____/\/_/ \/_/ 

Run any command or script as a daemon.

demon 
  help - show this usage
  spawn <command(s)> - create a new daemon
  destroy <pid|regex> - destroy a previously created daemon
  list <?pid|regex> - list running daemons

For issues https://github.com/streamwithme/demon.`
)

var demonCommands = map[string]func(...string){
	"spawn":   commands.Spawn,
	"destroy": commands.Destroy,
	"list":    commands.List,
}

func init() {
	logger.Formatter(func(level logger.Level, message string, vars ...interface{}) string {
		return fmt.Sprintf("demon > %s", fmt.Sprintf(message, vars...))
	})
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		for name, f := range demonCommands {
			if strings.EqualFold(name, args[0]) {
				f(args[1:]...)
				return
			}
		}
	}

	fmt.Println(usage)
}
