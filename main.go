package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"

	tables "github.com/tatsushid/go-prettytable"
)

const usage = ` _____   ______  __    __  ______  __   __    
/\  __-./\  ___\/\ "-./  \/\  __ \/\ "-.\ \   
\ \ \/\ \ \  __\\ \ \-./\ \ \ \/\ \ \ \-.  \  
 \ \____-\ \_____\ \_\ \ \_\ \_____\ \_\\"\_\ 
  \/____/ \/_____/\/_/  \/_/\/_____/\/_/ \/_/ 

Run any command or script as a daemon.

demon help 			show this usage
demon spawn <command(s)>	create a new daemon
demon kill <pid|regex>		kill a previously created daemon
demon list <?pid|regex>		list running daemons

Have a suggestion or want to report a bug?
https://github.com/streamwithme/demon/issues`

var args = os.Args[1:]

func daemons() []*process.Process {
	found := []*process.Process{}
	list, _ := ps.Processes()

	for _, proc := range list {
		if proc.Pid() != os.Getpid() && strings.Contains(proc.Executable(), "demon") {
			demon, _ := process.NewProcess(int32(proc.Pid()))
			found = append(found, demon)
		}
	}

	return found
}

func matches(proc *process.Process, expression string) bool {
	if pid, _ := strconv.Atoi(expression); pid == int(proc.Pid) {
		return true
	}

	command, _ := proc.Cmdline()

	if matched, _ := regexp.Match(expression, []byte(command)); matched {
		return true
	}

	return false
}

func spawn() {
	if len(args) < 1 {
		fmt.Println("No command or script provided.")
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

	proc, _ := process.NewProcess(int32(os.Getpid()))
	exe, _ := proc.Exe()
	cmd := exec.Command(exe, args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	go cmd.Run()

	for {
		if proc := cmd.Process; proc != nil {
			fmt.Printf("A new demon has been spawned with PID %d.\n", proc.Pid)
			proc.Release()
			break
		}
	}
}

func kill() {
	if len(args) < 1 {
		fmt.Println("No expression or PID was provided.")
		return
	}

	expression := args[0]
	killed := 0

	for _, demon := range daemons() {
		if matches(demon, expression) {
			if err := syscall.Kill(-int(demon.Pid), syscall.SIGKILL); err != nil {
				demon.Kill()
			}
			killed++
		}
	}

	if killed < 1 {
		fmt.Println("No demons found.")
		return
	}

	fmt.Printf("Killed %d demon(s).\n", killed)
}

func list() {
	var (
		expression = "."
		table, _   = tables.NewTable([]tables.Column{{Header: "PID"}, {Header: "Uptime"}, {Header: "Exec"}, {Header: "Command"}}...)
	)

	if len(args) > 0 {
		expression = args[0]
	}

	for _, demon := range daemons() {
		demons := []*process.Process{}
		cmd, _ := demon.CmdlineSlice()
		created, _ := demon.CreateTime()
		proc, _ := ps.FindProcess(int(demon.Pid))

		switch len(args) > 1 {
		case true && matches(demon, expression):
		case false:
			demons = append(demons, demon)
		}

		for _, demon := range demons {
			table.AddRow(demon.Pid, humanize.Time(time.Unix(created/1000, 0)), proc.Executable(), strings.Join(cmd[2:len(cmd)-1], " "))
		}
	}

	table.Separator = "	"
	table.Print()
}

func main() {
	if len(args) < 1 || strings.EqualFold(args[0], "help") {
		fmt.Println(usage)
		return
	}

	for label, f := range map[string]func(){
		"spawn":  spawn,
		"new":    spawn,
		"start":  spawn,
		"run":    spawn,
		"kill":   kill,
		"status": list,
		"list":   list,
	} {
		if strings.EqualFold(label, args[0]) {
			args = args[1:]
			f()
			return
		}
	}

	fmt.Println(usage)

}
