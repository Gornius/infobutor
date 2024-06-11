package manager

import (
	"testing"

	"github.com/gornius/infobutor/message"
	"github.com/gornius/infobutor/sender"
	"github.com/stretchr/testify/assert"
)

type MockSender struct {
	Config *MockSenderConfig
}

type MockSenderConfig struct {
	SomeParam string `json:"some_param"`
}

func (s *MockSender) LoadConfig(config map[string]any) error {
	return nil
}

func (s *MockSender) Send(message message.Message) error {
	return nil
}

func TestSenderFromConfig(t *testing.T) {
	assert := assert.New(t)
	manager := Manager{
		Factories: map[string]func() sender.Sender{
			"mock": func() sender.Sender {
				return new(MockSender)
			},
		},
	}
	sender, _ := manager.SenderFromConfig(map[string]any{
		"type": "mock",
	})
	assert.IsType(&MockSender{}, sender)
}

func TestRegister(t *testing.T) {
	assert := assert.New(t)
	manager := New()

	// Register first item
	err := manager.Register("mock", func() sender.Sender {
		return new(MockSender)
	})
	assert.Equal(nil, err)

	// Does mock key exist?
	_, ok := manager.Factories["mock"]
	assert.True(ok)

	// Try to register another sender with the same key
	err = manager.Register("mock", func() sender.Sender {
		return new(MockSender)
	})
	assert.NotEqual(nil, err)
}
