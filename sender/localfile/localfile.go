package localfile

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gornius/infobutor/message"
	"github.com/gornius/infobutor/pkg/pathutils"
)

type LocalFileSender struct {
	Config *LocalFileSenderConfig
}

type LocalFileSenderConfig struct {
	Path      string `json:"path"`
	SplitDays bool   `json:"split_days"`
}

func parsePath(path string, splitDays bool) (string, error) {
	dir := filepath.Dir(path)
	ext := filepath.Ext(path)
	base := strings.TrimSuffix(filepath.Base(path), ext)

	dir, err := pathutils.ExpandTildeToHomeDir(dir)
	if err != nil {
		return "", err
	}

	if splitDays {
		fmt.Println(base)
		base += "_" + time.Now().Format(time.DateOnly)
	}

	return filepath.Join(dir, base) + ext, nil
}

func (lfs *LocalFileSender) Send(message message.Message) error {
	filePath, err := parsePath(lfs.Config.Path, lfs.Config.SplitDays)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}
	defer file.Close()

	formattedMessage := "----------------------------------------\n"
	formattedMessage += " ORIGIN: " + message.Origin + "\n"
	formattedMessage += " TITLE: " + message.Title + "\n"
	formattedMessage += " DATE AND TIME: " + time.Now().Format(time.DateTime) + "\n"
	formattedMessage += "----------------------------------------\n"
	formattedMessage += message.Content + "\n\n"

	if _, err := file.WriteString(formattedMessage); err != nil {
		return err
	}

	return nil
}

func (lfs *LocalFileSender) LoadConfig(config map[string]any) error {
	lfs.Config = &LocalFileSenderConfig{}

	tempJson, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tempJson, lfs.Config)
	if err != nil {
		return err
	}

	return nil
}
