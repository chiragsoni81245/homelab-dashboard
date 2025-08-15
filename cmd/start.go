package cmd

import (
	"homelab-dashboard/internal/logger"
	"homelab-dashboard/internal/server"

	"github.com/spf13/cobra"
)


var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Its use to start the web UI server",
	Long: `Its use to start the web UI server`,
	Run: func(cmd *cobra.Command, args []string) { 
		// We need to call the server start function

		server := server.NewServer()

		err := server.Start()
		if err != nil {
			logger.Log.Fatal(err)
		}
	},
}
