package localfile

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/gornius/infobutor/message"
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
	base := filepath.Base(path)
	ext := filepath.Ext(path)

	fmt.Printf("dir: %v\n", dir)
	fmt.Printf("base: %v\n", base)
	fmt.Printf("ext: %v\n", ext)

	if dir[:1] == "~" {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		dir = filepath.Join(usr.HomeDir, dir[1:])
	}

	if splitDays {
		base += "_" + time.Now().Format(time.DateOnly)
	}

	return filepath.Join(dir, base) + "." + ext, nil
}

func (lfs *LocalFileSender) Send(message message.Message) error {
	filePath, err := parsePath(lfs.Config.Path, lfs.Config.SplitDays)
	fmt.Printf("filePath: %v\n", filePath)
	if err != nil {
		return err
	}
	if lfs.Config.SplitDays {
		filePath += "_" + time.Now().Format(time.DateOnly)
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
