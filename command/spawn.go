package command

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/takama/daemon"
)

func init() {
	GetCommand().AddCommand(&cobra.Command{
		Use:   "spawn",
		Short: "Daemonize a command or script.",
		Long:  "Daemonize a command or script.",
		Run: func(cmd *cobra.Command, args []string) {
			service, err := daemon.New("name", "description")
			if err != nil {
				log.Fatal("Error: ", err)
			}
			status, err := service.Install()
			if err != nil {
				log.Fatal(status, "\nError: ", err)
			}
			fmt.Println(status)
		},
	})
}
