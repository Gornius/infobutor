package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/gornius/infobutor/channel"
	"github.com/gornius/infobutor/sender"
	"github.com/gornius/infobutor/sender/manager"
)

type Config struct {
	Channels map[string]*channel.Channel
	Senders  map[string]sender.Sender
}

type ConfigParser interface {
	FromFile(manager *manager.Manager, file []byte) (*Config, error)
	FromMap(senderManager *manager.Manager, configMap map[string]any) (*Config, error)
	CreateDefault(path string) error
}

func DefaultLocation() string {
	return xdg.ConfigHome + "/infobutor/conf.yaml"
}

func FromFile(manager *manager.Manager, path string) (*Config, error) {
	var parser ConfigParser
	switch extension := filepath.Ext(path); extension {
	case ".yml", ".yaml":
		parser = &YamlParser{}
	default:
		return nil, errors.New("unsupported config type")
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err := parser.CreateDefault(path); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return parser.FromFile(manager, file)
}
