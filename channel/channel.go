package channel

import (
	"github.com/gornius/infobutor/message"
	"github.com/gornius/infobutor/sender"
)

type Channel struct {
	Name    string
	Token   string
	Senders []sender.Sender
}

func (ch *Channel) Send(message *message.Message) {
	for _, sender := range ch.Senders {
		sender.Send(*message)
	}
}
