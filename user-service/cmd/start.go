package cmd

import (
	"user-service/internal/app"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the user service",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunServer() //run server from app package
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
