package command

import "github.com/spf13/cobra"

func init() {
	base.AddCommand(&cobra.Command{
		Use:   "spawn",
		Short: "Daemonize a command or script.",
		Long:  "Daemonize a command or script.",
		Run: func(cmd *cobra.Command, args []string) {
			// woof
		},
	})
}
