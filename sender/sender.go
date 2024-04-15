package sender

import "github.com/gornius/infobutor/message"

type SenderConfig any

type Sender interface {
	Send(message message.Message) error
	LoadConfig(config map[string]any) error
}
