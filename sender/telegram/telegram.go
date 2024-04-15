package telegram

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gornius/infobutor/message"
)

type TelegramSender struct {
	Config *TelegramSenderConfig
}

type TelegramSenderConfig struct {
	ApiToken string
	ChatId   string
}

func (ts *TelegramSender) Send(message message.Message) error {
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

func (ts *TelegramSender) LoadConfig(config map[string]any) error {
	ts.Config = &TelegramSenderConfig{}
	ts.Config.ApiToken = config["api_token"].(string)
	chatId, ok := config["chat_id"].(string)
	if !ok {
		chatId = strconv.Itoa(config["chat_id"].(int))
	}
	ts.Config.ChatId = chatId

	return nil
}
