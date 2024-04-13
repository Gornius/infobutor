package channel

import "github.com/gornius/infobutor/sender"

type Channel struct {
	Name   string
	Token  string
	Sender *sender.Sender
}
