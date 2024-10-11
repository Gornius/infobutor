package manager

import (
	"errors"
	"fmt"

	"github.com/gornius/infobutor/sender"
	"github.com/gornius/infobutor/sender/discord"
	"github.com/gornius/infobutor/sender/localfile"
	"github.com/gornius/infobutor/sender/telegram"
)

// manages senders types available in application
type Manager struct {
	Factories map[string]func() sender.Sender
}

type Option func(m *Manager) error

// Option that includes built-in senders
func GetBuiltInSendersFactories() map[string]func() sender.Sender {
	return map[string]func() sender.Sender{
		"discord": func() sender.Sender {
			return new(discord.DiscordWebhookSender)
		},
		"telegram": func() sender.Sender {
			return new(telegram.TelegramSender)
		},
		"localfile": func() sender.Sender {
			return new(localfile.LocalFileSender)
		},
	}
}

// tries to create a sender from map[string]any
func (manager *Manager) SenderFromConfig(config map[string]any) (sender.Sender, error) {
	senderKey := config["type"].(string)
	sender := manager.Factories[senderKey]()
	sender.LoadConfig(config)

	return sender, nil
}

// adds a sender to available sender types
func (manager *Manager) Register(name string, factory func() sender.Sender) error {
	if manager.Factories[name] != nil {
		return errors.New("tried to register a sender with a name that's already registered")
	}
	fmt.Printf("manager: %+v\n", manager)
	manager.Factories[name] = factory
	return nil
}

func (manager *Manager) RegisterBuiltInSenders() error {
	for name, factory := range GetBuiltInSendersFactories() {
		err := manager.Register(name, factory)
		if err != nil {
			return err
		}
	}
	return nil
}

// creates new sender manager instance
func New(opts ...Option) (*Manager, error) {
	manager := &Manager{}
	manager.Factories = map[string]func() sender.Sender{}
	for _, opt := range opts {
		err := opt(manager)
		if err != nil {
			return nil, err
		}
	}
	return manager, nil
}
