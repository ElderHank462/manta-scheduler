package util

import (
	"os"
	"path/filepath"
)

func GetTasksPath() string {
	home, _ := os.UserHomeDir()
	folderpath := filepath.Join(home, ".manta")

	os.MkdirAll(folderpath, 0755)

	return filepath.Join(folderpath, "tasks.json")
}