package cmd

import (
	"homelab-dashboard/internal/config"
	"homelab-dashboard/internal/database"
	"homelab-dashboard/internal/logger"
	"os"

	"github.com/spf13/cobra"
)

var configPath string


var rootCmd = &cobra.Command{
	Use:   "homelab-dashboard",
	Short: "Its a CLI application which will be used to host a web UI for your machine stats",
	Long: `homelab-dashboard CLI application will provide a way to run a web UI for your machine stats`,
	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	cobra.OnInitialize(initApp)
	err := rootCmd.Execute()
	if err != nil {
		logger.Log.Error(err)
		os.Exit(1)
	}
}

func initApp() {
	logger.InitLogger()
	config.LoadConfig(configPath)
	database.InitDB(config.App.DatabaseURI)
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "config file")

	rootCmd.AddCommand(startCmd)
}


