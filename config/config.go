package config

import (
	"github.com/adrg/xdg"
	"github.com/gornius/infobutor/channel"
	"github.com/gornius/infobutor/sender"
)

type Config struct {
	Channels map[string]*channel.Channel
	Senders  map[string]sender.Sender
}

func DefaultLocation() string {
	return xdg.ConfigHome + "/infobutor/conf.yaml"
}
