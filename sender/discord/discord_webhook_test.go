package discord

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	assert := assert.New(t)
	url := "http://test-webhoo.com"
	configMap := map[string]any{
		"webhook_url": url,
	}
	expected := &DiscordWebhookSenderConfig{
		WebhookUrl: url,
	}
	s := DiscordWebhookSender{}
	s.LoadConfig(configMap)
	got := s.Config
	assert.EqualValues(expected, got)
}
