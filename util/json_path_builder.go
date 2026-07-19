/*
Utility functions to build paths to the tasks.json file and the taskgroups.json file.
*/

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

func GetGroupsPath() string {
	home, _ := os.UserHomeDir()
	folderpath := filepath.Join(home, ".manta")

	os.MkdirAll(folderpath, 0755)

	return filepath.Join(folderpath, "groups.json")
}