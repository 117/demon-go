package command

import "github.com/spf13/cobra"

func init() {
	GetCommand().AddCommand(&cobra.Command{
		Use:   "destroy",
		Short: "Kill a specified daemon.",
		Long:  "Kill a specified daemon.",
		Run: func(cmd *cobra.Command, args []string) {
			// woof
		},
	})
}
