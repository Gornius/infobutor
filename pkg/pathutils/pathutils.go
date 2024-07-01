package pathutils

import (
	"os"
	"path/filepath"
	"strings"
)

// takes a path with optional "~/" at the beginning and expands it to current user's home folder
func ExpandTildeToHomeDir(path string) (string, error) {
	currentUserHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if path == "~" {
		return currentUserHome, nil
	}

	if strings.HasPrefix(path, "~/") {
		return filepath.Join(currentUserHome, path[2:]), nil
	}

	return path, nil
}
