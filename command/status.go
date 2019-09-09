package command

import "github.com/spf13/cobra"

func init() {
	base.AddCommand(&cobra.Command{
		Use:   "status",
		Short: "Check the uptime and recent output of a daemon.",
		Long:  "Check the uptime and recent output of a daemon.",
		Run: func(cmd *cobra.Command, args []string) {
			// woof
		},
	})
}
