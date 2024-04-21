package config

import (
	_ "embed"
	"os"

	"github.com/gornius/infobutor/channel"
	"github.com/gornius/infobutor/sender"
	"github.com/gornius/infobutor/sender/manager"
	"gopkg.in/yaml.v3"
)

type YamlParser struct {
}

func (yp *YamlParser) FromFile(senderManager *manager.Manager, file []byte) (*Config, error) {
	configMap := map[string]any{}

	if err := yaml.Unmarshal(file, configMap); err != nil {
		return nil, err
	}
	return yp.FromMap(senderManager, configMap)
}

// TODO: move all of this to config - no reason to make this yaml specific
func (yp *YamlParser) FromMap(senderManager *manager.Manager, configMap map[string]any) (*Config, error) {
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

//go:embed default_configs/conf.yml
var defaultYaml []byte

func (yp *YamlParser) CreateDefault(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	file.Write(defaultYaml)
	return nil
}
