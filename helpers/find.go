package helpers

import (
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"
)

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
