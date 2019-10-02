package main

import (
	"os"
	"os/exec"

	"github.com/117/logger"
	"github.com/shirou/gopsutil/process"
)

// Spawn creates a new daemon with the provided command.
func Spawn(args ...string) {
	if len(args) < 1 {
		logger.Log(logger.Err, "No command or script provided.")
		return
	}

	// This is marked as a demon.
	if args[len(args)-1] == "--demon" {
		exec.Command(args[0], args[1:len(args)-1]...).Run()
		return
	}

	// Re-run this command but with the --demon flag.
	args = append([]string{"spawn"}, args...)
	args = append(args, "--demon")

	var (
		proc, _ = process.NewProcess(int32(os.Getpid()))
		exe, _  = proc.Exe()
		cmd     = exec.Command(exe, args...)
	)

	go cmd.Run()

	for {
		if proc := cmd.Process; proc != nil {
			logger.Logf(logger.Info, "A new demon has been spawned with PID %d.", proc.Pid)
			proc.Release()
			break
		}
	}
}
