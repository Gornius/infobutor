package localfile

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParsePath(t *testing.T) {
	assert := assert.New(t)
	currentDay := time.Now().Format(time.DateOnly)
	var path, expected, got string

	path = "/var/run/test/testfile.txt"
	expected = path
	got, _ = parsePath(path, false)
	assert.Equal(expected, got, "Parse path without splitting days")

	path = "/var/run/test/testfile.txt"
	expected = "/var/run/test/testfile_" + currentDay + ".txt"
	got, _ = parsePath(path, true)
	assert.Equal(expected, got, "Parse path with splitting days")
}

func TestLoadConfig(t *testing.T) {
	assert := assert.New(t)
	path := "~/.local/infobutor.log"
	splitDays := false
	configMap := map[string]any{
		"path":       path,
		"split_days": splitDays,
	}
	expected := &LocalFileSenderConfig{
		Path:      path,
		SplitDays: splitDays,
	}
	lfs := LocalFileSender{}
	lfs.LoadConfig(configMap)
	got := lfs.Config
	assert.EqualValues(expected, got)
}
