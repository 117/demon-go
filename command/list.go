package command

import (
	"github.com/spf13/cobra"
)

func init() {
	base.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "View a list of active daemons.",
		Long:  "View a list of active daemons.",
		Run: func(cmd *cobra.Command, args []string) {
			// woof
		},
	})
}
