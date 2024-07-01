package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/gornius/infobutor/pkg/randstring"
	"github.com/gornius/infobutor/sender"
	"github.com/gornius/infobutor/sender/manager"
	"github.com/gornius/infobutor/sink"
)

type Config struct {
	Secret     string
	ConfigPath string
	Sinks   map[string]*sink.Sink
	Senders    map[string]sender.Sender
}

type ConfigParser interface {
	FromFile(bytes *[]byte) (map[string]any, error)
	ToFile(configMap map[string]any) (*[]byte, error)
}

func DefaultLocation() string {
	return xdg.ConfigHome + "/infobutor/conf.yaml"
}

func createDefaultConfigFileIfDoesNotExist(path string, parser ConfigParser) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		bytes, err := parser.ToFile(GetDefaultConfig())
		if err != nil {
			return err
		}
		file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		file.Write(*bytes)
	} else if err != nil {
		return err
	}
	return nil
}

func FromFile(manager *manager.Manager, path string) (*Config, error) {
	var parser ConfigParser
	switch extension := filepath.Ext(path); extension {
	case ".yml", ".yaml":
		parser = &YamlParser{}
	default:
		return nil, errors.New("unsupported config type")
	}

	createDefaultConfigFileIfDoesNotExist(path, parser)

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	configMap, err := parser.FromFile(&file)
	if err != nil {
		return nil, err
	}

	return FromMap(manager, configMap)
}

type configStructure struct {
	Secret   string                    `json:"secret"`
	Senders  map[string]map[string]any `json:"senders"`
	Sinks map[string]struct {
		Name    string   `json:"name"`
		Token   string   `json:"token"`
		Senders []string `json:"senders"`
	} `json:"sinks"`
}

func FromMap(senderManager *manager.Manager, configMap map[string]any) (*Config, error) {
	tempJson, err := json.Marshal(configMap)
	if err != nil {
		return nil, err
	}

	cfg := &configStructure{}
	err = json.Unmarshal(tempJson, cfg)
	if err != nil {
		return nil, err
	}

	senders := map[string]sender.Sender{}
	sendersConfig := cfg.Senders
	for senderId, senderConfig := range sendersConfig {
		newSender, err := senderManager.SenderFromConfig(senderConfig)
		if err != nil {
			return nil, err
		}
		senders[senderId] = newSender
	}

	sinks := map[string]*sink.Sink{}
	for sinkId, sinkConfig := range cfg.Sinks {
		sinkSenders := []sender.Sender{}
		for _, senderId := range sinkConfig.Senders {
			sender, ok := senders[senderId]
			if !ok {
				return nil, errors.New("sink tried to use sender that doesn't exist")
			}
			sinkSenders = append(sinkSenders, sender)
		}
		sinks[sinkId] = &sink.Sink{
			Name:    sinkConfig.Name,
			Token:   sinkConfig.Token,
			Senders: sinkSenders,
		}
	}

	config := new(Config)
	config.Senders = senders
	config.Sinks = sinks
	config.Secret = cfg.Secret
	return config, nil
}

func GetDefaultConfig() map[string]any {
	return map[string]any{
		"secret": randstring.WithLength(20),
		"senders": map[string]any{
			"my_telegram": map[string]any{
				"type":      "telegram",
				"api_token": "YOUR_API_TOKEN",
				"chat_id":   "-123456789",
			},
		},
		"sinks": map[string]any{
			"default": map[string]any{
				"token":   "2018s1m0eme0i1we21ps21",
				"senders": []string{"my_telegram"},
			},
		},
	}
}
