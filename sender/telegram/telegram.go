package telegram

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gornius/infobutor/message"
)

type telegramSender struct {
	Config *TelegramSenderConfig
}

type TelegramSenderConfig struct {
	ApiToken string
	ChatId   string
}

func New(config TelegramSenderConfig) *telegramSender {
	return &telegramSender{Config: &config}
}

func (ts *telegramSender) Send(message *message.Message, channelName string) error {
	endpoint, err := url.Parse("http://api.telegram.org/bot" + ts.Config.ApiToken + "/sendMessage")

	if err != nil {
		return err
	}

	query := endpoint.Query()
	query.Set("chat_id", ts.Config.ChatId)
	query.Set("text", message.Content)
	endpoint.RawQuery = query.Encode()

	result, err := http.Post(endpoint.String(), "none", nil)
	if err != nil {
		return err
	}

	if result.StatusCode != http.StatusOK {
		resultBody, err := io.ReadAll(result.Body)
		defer result.Body.Close()

		if err != nil {
			return err
		}

		fmt.Printf("resultBody: %v\n", string(resultBody))
	}

	return nil
}
