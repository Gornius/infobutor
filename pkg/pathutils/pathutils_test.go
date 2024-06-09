package pathutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExpandTildeToHomeDir(t *testing.T) {
	var path, expanded, expected string

	userHomeDir, _ := os.UserHomeDir()

	path = "~"
	expanded, _ = ExpandTildeToHomeDir(path)
	expected = userHomeDir
	if expanded != expected {
		t.Errorf("Failed to expand %s to expected location %s, got: %s", path, expected, expanded)
	}

	path = filepath.Join("~", ".local", "bin")
	expanded, _ = ExpandTildeToHomeDir(path)
	expected = filepath.Join(userHomeDir, ".local", "bin")
	if expanded != expected {
		t.Errorf("Failed to expand %s to expected location %s, got: %s", path, expected, expanded)
	}

	path = filepath.Join("/", "usr", "lib")
	expected = path
	expanded, _ = ExpandTildeToHomeDir(path)
	if expanded != expected {
		t.Errorf("Failed to expand %s to expected location %s, got: %s", path, expected, expanded)
	}

}
