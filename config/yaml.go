package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gornius/infobutor/channel"
	"github.com/gornius/infobutor/sender"
	"github.com/gornius/infobutor/sender/manager"
	"gopkg.in/yaml.v3"
)

func FromFile(senderManager *manager.Manager, filePath string) (*Config, error) {
	configMap := map[string]any{}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	switch extension := filepath.Ext(filePath); extension {
	case ".yml", ".yaml":
		err = yaml.Unmarshal(file, configMap)
		if err != nil {
			return nil, err
		}
		return FromMap(senderManager, configMap)
	}

	return nil, errors.New("unsupported config file type")

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
			channels[channelId] = &channel.Channel{
				Token:   channelConfig.(map[string]any)["token"].(string),
				Senders: []sender.Sender{senders[channelConfig.(map[string]any)["sender"].(string)]},
			}
		}
	}

	config := new(Config)

	config.Senders = senders
	config.Channels = channels

	return config, nil

}
