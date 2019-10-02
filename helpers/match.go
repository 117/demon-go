package helpers

import (
	"fmt"
	"regexp"

	"github.com/shirou/gopsutil/process"
)

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
