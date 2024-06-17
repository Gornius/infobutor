package manager

import (
	"errors"

	"github.com/gornius/infobutor/sender"
	"github.com/gornius/infobutor/sender/discord"
	"github.com/gornius/infobutor/sender/localfile"
	"github.com/gornius/infobutor/sender/telegram"
)

type Manager struct {
	Factories map[string]func() sender.Sender
}

func NewWithAllBuiltIn() *Manager {
	return &Manager{
		Factories: map[string]func() sender.Sender{
			"discord": func() sender.Sender {
				return new(discord.DiscordWebhookSender)
			},
			"telegram": func() sender.Sender {
				return new(telegram.TelegramSender)
			},
			"localfile": func() sender.Sender {
				return new(localfile.LocalFileSender)
			},
		},
	}
}

func (manager *Manager) SenderFromConfig(config map[string]any) (sender.Sender, error) {
	senderKey := config["type"].(string)
	sender := manager.Factories[senderKey]()
	sender.LoadConfig(config)

	return sender, nil
}

func (manager *Manager) Register(name string, factory func() sender.Sender) error {
	if manager.Factories[name] != nil {
		return errors.New("tried to register a sender with a name that's already registered")
	}
	manager.Factories[name] = factory
	return nil
}

func New() *Manager {
	return &Manager{
		Factories: map[string]func() sender.Sender{},
	}
}
