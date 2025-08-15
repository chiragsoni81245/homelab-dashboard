package server

import (
	"homelab-dashboard/internal/logger"
	"os"
	"path"
)

func GetTemplateFilePaths(baseDir string) []string {
	files, err := os.ReadDir(baseDir)
	if err != nil {
		logger.Log.Fatal(err)
	}

	var templateFiles []string
	for _, file := range files {
		if !file.IsDir() {
			templateFiles = append(templateFiles, path.Join(baseDir, file.Name()))
		}
	}

	return templateFiles
}
