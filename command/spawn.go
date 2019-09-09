package command

import (
	"strings"

	"github.com/117/logger"

	"../service"
	"github.com/spf13/cobra"
)

func init() {
	GetCommand().AddCommand(&cobra.Command{
		Use:   "spawn",
		Short: "Daemonize a command or script.",
		Long:  "Daemonize a command or script.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := new(service.Service).Spawn(); err != nil {
				logger.Log(logger.Fatal, strings.ToLower(err.Error()))
			}
		},
	})
}
