package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/117/logger"
	"github.com/dustin/go-humanize"
	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"
	tables "github.com/tatsushid/go-prettytable"
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

type Demon struct {
	*process.Process
}

func (d *Demon) Match(expr string) bool {
	// Check if the PID string val matches.
	if matched, _ := regexp.Match(expr, []byte(fmt.Sprintf("%d", d.Pid))); matched {
		return true
	}

	var command, _ = d.Cmdline()

	if matched, _ := regexp.Match(expr, []byte(command)); matched {
		return true
	}

	return false
}

func demons() []*Demon {
	var (
		found   []*Demon
		list, _ = ps.Processes()
	)

	for _, proc := range list {
		if proc.Pid() == os.Getpid() {
			continue
		}

		if strings.Contains(proc.Executable(), "demon") {
			demon, _ := process.NewProcess(int32(proc.Pid()))
			found = append(found, &Demon{demon})
		}
	}

	return found
}

func init() {
	logger.Formatter(func(level logger.Level, message string, vars ...interface{}) string {
		return fmt.Sprintf("demon > %s", fmt.Sprintf(message, vars...))
	})
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		logger.Log(logger.Info, "usage")
		fmt.Println(usage)
		return
	}

	switch strings.ToLower(args[0]) {
	case "spawn":
		if len(args) < 2 {
			logger.Log(logger.Err, "No command or script provided.")
			return
		}

		dargs := args[1:len(args)]

		// Re-run this command but with the --demon flag.
		if args[len(args)-1] != "--demon" {
			var (
				proc, _ = process.NewProcess(int32(os.Getpid()))
				exe, _  = proc.Exe()
				cmd     = exec.Command(exe, "spawn", strings.Join(dargs, " "), "--demon")
			)

			go cmd.Run()

			for {
				if proc := cmd.Process; proc != nil {
					logger.Logf(logger.Info, "A new demon has been spawned with PID %d.", proc.Pid)
					proc.Release()
					break
				}
			}

			return
		}

		parts := strings.Split(dargs[0], " ")
		cmd := exec.Command(parts[0], parts[1:]...)

		cmd.Run()
		return
	case "destroy":
		if len(args) < 2 {
			logger.Log(logger.Err, "No expression or PID was provided.")
			return
		}

		var killed int

		for _, demon := range demons() {
			if demon.Match(args[1]) {
				demon.Kill()
				killed++
			}
		}

		logger.Logf(logger.Info, "Found and destroyed %d demon(s).", killed)
		return
	case "list":
		logger.Logf(logger.Info, "list")
		table, _ := tables.NewTable(
			[]tables.Column{
				{Header: "PID"},
				{Header: "Uptime"},
				{Header: "Exec"},
				{Header: "Command"},
			}...)

		table.Separator = "	"

		for _, demon := range demons() {
			var (
				demons     = []*Demon{}
				cmd, _     = demon.CmdlineSlice()
				created, _ = demon.CreateTime()
				proc, _    = ps.FindProcess(int(demon.Pid))
			)

			switch len(args) > 1 {
			case true:
				if demon.Match(args[0]) {
					demons = append(demons, demon)
				}
			case false:
				demons = append(demons, demon)
			}

			var (
				uptime  = humanize.Time(time.Unix(created/1000, 0))
				command = strings.Join(cmd[2:len(cmd)-1], " ")
			)

			for _, demon := range demons {
				table.AddRow(demon.Pid, uptime, proc.Executable(), command)
			}
		}

		table.Print()
		return
	}

	logger.Log(logger.Info, "usage")
	fmt.Println(usage)
}
