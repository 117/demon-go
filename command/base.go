package command

import (
	"log"

	"github.com/spf13/cobra"
)

var base = &cobra.Command{
	Use:   "demon",
	Short: "Daemonize and manage scripts or commands.",
	Long:  "Daemonize and manage scripts or commands.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("demon command")
	},
}

// GetBase - no description
func GetBase() *cobra.Command {
	return base
}
