package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gornius/infobutor/message"
)

// sender that uses telegram bot to send messages
type TelegramSender struct {
	Config *TelegramSenderConfig
}

type TelegramSenderConfig struct {
	ApiToken string `json:"api_token"`
	ChatId   int    `json:"chat_id"`
}

func (ts *TelegramSender) Send(message message.Message) error {
	endpoint, err := url.Parse("http://api.telegram.org/bot" + ts.Config.ApiToken + "/sendMessage")

	if err != nil {
		return err
	}

	query := endpoint.Query()
	query.Set("chat_id", strconv.Itoa(ts.Config.ChatId))
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
	tempJson, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tempJson, ts.Config)
	if err != nil {
		return err
	}

	return nil
}
