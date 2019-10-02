package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/117/logger"
	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"
)

const usage = ` _____   ______  __    __  ______  __   __    
/\  __-./\  ___\/\ "-./  \/\  __ \/\ "-.\ \   
\ \ \/\ \ \  __\\ \ \-./\ \ \ \/\ \ \ \-.  \  
 \ \____-\ \_____\ \_\ \ \_\ \_____\ \_\\"\_\ 
  \/____/ \/_____/\/_/  \/_/\/_____/\/_/ \/_/ 

Run any command or script as a daemon.

demon
	help 			show this usage
	spawn <command(s)>	create a new daemon
	destroy <pid|regex>	destroy a previously created daemon
	list <?pid|regex>	ist running daemons

Have a suggestion? (https://github.com/streamwithme/demon/issues)`

// FindMatchingDemons with matching pid or command.
func FindMatchingDemons(expression string) []*process.Process {
	var (
		found   []*process.Process
		list, _ = ps.Processes()
	)

	for _, proc := range list {
		if proc.Pid() == os.Getpid() {
			continue
		}

		if strings.Contains(proc.Executable(), "demon") {
			demon, _ := process.NewProcess(int32(proc.Pid()))
			found = append(found, demon)
		}
	}

	return found
}

// MatchDemon against pid and command.
func MatchDemon(proc *process.Process, expression string) bool {
	// Check if the PID string val matches.
	if matched, _ := regexp.Match(expression, []byte(fmt.Sprintf("%d", proc.Pid))); matched {
		return true
	}

	var command, _ = proc.Cmdline()

	if matched, _ := regexp.Match(expression, []byte(command)); matched {
		return true
	}

	return false
}

func init() {
	logger.Formatter(func(level logger.Level, message string, vars ...interface{}) string {
		return fmt.Sprintf("demon > %s", fmt.Sprintf(message, vars...))
	})
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 || strings.EqualFold(args[0], "help") {
		fmt.Println(usage)
		return
	}

	switch strings.ToLower(args[0]) {
	case "s":
	case "spawn":
		Spawn(args[1:]...)
		return
	case "d":
	case "kill":
	case "destroy":
		Destroy(args[1:]...)
		return
	case "l":
	case "list":
		List(args[1:]...)
		return
	}

	fmt.Println(usage)
}
