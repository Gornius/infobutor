package manager

import (
	"errors"
	"reflect"

	"github.com/gornius/infobutor/sender"
	"github.com/gornius/infobutor/sender/discord"
	"github.com/gornius/infobutor/sender/telegram"
)

type Manager struct {
	Senders map[string]reflect.Type
}

func WithAllBuiltIn() Manager {
	return Manager{
		Senders: map[string]reflect.Type{
			"telegram": reflect.TypeOf(telegram.TelegramSender{}),
			"discord":  reflect.TypeOf(discord.DiscordWebhookSender{}),
		},
	}
}

func (manager *Manager) SenderFromConfig(config map[string]any) (sender.Sender, error) {
	sender := reflect.New(manager.Senders[config["type"].(string)]).Interface().(sender.Sender)
	sender.LoadConfig(config)

	return sender, nil
}

func (manager *Manager) Register(name string, sender sender.Sender) error {
	if manager.Senders[name] != nil {
		return errors.New("tried to register a sender with a name that's already registered")
	}
	manager.Senders[name] = reflect.TypeOf(sender)
	return nil
}
