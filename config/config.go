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
	FromFile(manager *manager.Manager, file []byte) (map[string]any, error)
	ToFile(path string, configMap map[string]any) error
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
		if err := parser.ToFile(path, GetDefaultConfig()); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	configMap, err := parser.FromFile(manager, file)
	if err != nil {
		return nil, err
	}

	return FromMap(manager, configMap)
}

func FromMap(senderManager *manager.Manager, configMap map[string]any) (*Config, error) {
	senders := map[string]sender.Sender{}
	if configMap["channels"] != nil {
		sendersConfigMap := configMap["senders"].(map[string]any)
		for senderIdentifier, senderConfig := range sendersConfigMap {
			newSender, err := senderManager.SenderFromConfig(senderConfig.(map[string]any))
			if err != nil {
				return nil, err
			}
			senders[senderIdentifier] = newSender
		}
	}

	channels := map[string]*channel.Channel{}
	if configMap["channels"] != nil {
		channelConfigMap := configMap["channels"].(map[string]any)
		for channelId, channelConfig := range channelConfigMap {
			channelSenders := []sender.Sender{}
			for _, senderId := range channelConfig.(map[string]any)["senders"].([]any) {
				channelSenders = append(channelSenders, senders[senderId.(string)])
			}
			channels[channelId] = &channel.Channel{
				Token:   channelConfig.(map[string]any)["token"].(string),
				Senders: channelSenders,
			}
		}
	}

	config := new(Config)

	config.Senders = senders
	config.Channels = channels

	return config, nil
}

func GetDefaultConfig() map[string]any {
	return map[string]any{
		"senders": map[string]any{
			"my_telegram": map[string]any{
				"type":      "telegram",
				"api_token": "YOUR_API_TOKEN",
				"chat_id":   "-123456789",
			},
		},
		"channels": map[string]any{
			"default": map[string]any{
				"token":   "2018s1m0eme0i1we21ps21",
				"senders": []string{"my_telegram"},
			},
		},
	}
}
