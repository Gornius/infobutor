package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gornius/infobutor/message"
)

type DiscordWebhookSender struct {
	Config *DiscordWebhookSenderConfig
}

type DiscordWebhookSenderConfig struct {
	WebhookUrl string `json:"webhook_url"`
}

func (s *DiscordWebhookSender) Send(message message.Message) error {
	endpoint, err := url.Parse(s.Config.WebhookUrl)
	if err != nil {
		return err
	}

	requestBody, err := json.Marshal(map[string]any{
		"content": message.Content,
	})
	if err != nil {
		return err
	}

	result, err := http.Post(endpoint.String(), "application/json", bytes.NewReader(requestBody))
	if err != nil {
		return err
	}

	if result.StatusCode != http.StatusNoContent {
		defer result.Body.Close()
		resultBody, err := io.ReadAll(result.Body)

		if err != nil {
			return err
		}

		fmt.Printf("resultBody: %v\n", string(resultBody))
	}

	return nil
}

func (s *DiscordWebhookSender) LoadConfig(config map[string]any) error {
	s.Config = &DiscordWebhookSenderConfig{}
	tempJson, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tempJson, s.Config)
	if err != nil {
		return err
	}

	return nil
}
