package main

import (
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"

	tables "github.com/tatsushid/go-prettytable"
)

// List returns a list of running daemons.
func List(args ...string) {
	var (
		expression = "*"
		table, _   = tables.NewTable([]tables.Column{{Header: "PID"}, {Header: "Uptime"}, {Header: "Exec"}, {Header: "Command"}}...)
	)

	if len(args) > 0 {
		expression = args[0]
	}

	for _, demon := range FindMatchingDemons(expression) {
		var (
			demons     = []*process.Process{}
			cmd, _     = demon.CmdlineSlice()
			created, _ = demon.CreateTime()
			proc, _    = ps.FindProcess(int(demon.Pid))
		)

		switch len(args) > 1 {
		case true:
			if MatchDemon(demon, expression) {
				demons = append(demons, demon)
			}
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
