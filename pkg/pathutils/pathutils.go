package pathutils

import (
	"os"
	"path/filepath"
	"strings"
)

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
