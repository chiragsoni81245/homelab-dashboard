package server

import (
	"os"
	"path"
)

func GetTemplateFile(baseDir string) ([]string, error) {
	files, err := os.ReadDir(baseDir)
	if err != nil {
		return nil, err
	}

	var templateFiles []string
	for _, file := range files {
		if !file.IsDir() {
			templateFiles = append(templateFiles, path.Join(baseDir, file.Name()))
		}
	}

	return templateFiles, nil
}
