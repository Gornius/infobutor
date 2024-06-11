package telegram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	assert := assert.New(t)
	apiToken := "test123"
	chatId := -1

	configMap := map[string]any{
		"api_token": apiToken,
		"chat_id":   chatId,
	}

	expected := &TelegramSenderConfig{
		ApiToken: apiToken,
		ChatId:   chatId,
	}
	s := TelegramSender{}
	s.LoadConfig(configMap)
	got := s.Config
	assert.EqualValues(expected, got)
}
