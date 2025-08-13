package main

import (
	"homelab-dashboard/cmd"
	"homelab-dashboard/internal/logger"
)

func main() {
	logger.InitLogger()
	cmd.Execute()
}
