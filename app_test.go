package infobutor

import (
	"testing"

	"github.com/gornius/infobutor/channel"
	"github.com/gornius/infobutor/config"
	"github.com/gornius/infobutor/sender"
	"github.com/stretchr/testify/assert"
)

func TestGetChannelByToken(t *testing.T) {
	assert := assert.New(t)

	app := NewApp()
	app.Config = &config.Config{
		Channels: map[string]*channel.Channel{
			"test1_id": {
				Name:    "test1_name",
				Token:   "test1_token",
				Senders: []sender.Sender{},
			},
			"test2_id": {
				Name:    "test2_name",
				Token:   "test2_token",
				Senders: []sender.Sender{},
			},
		},
	}

	var (
		ch  *channel.Channel
		err error
	)

	ch, err = app.GetChannelByToken("test1_token")
	assert.Nil(
		err,
	)
	assert.Equal(
		"test1_name",
		ch.Name,
	)

	ch, err = app.GetChannelByToken("test2_token")
	assert.Nil(
		err,
	)
	assert.Equal(
		"test2_name",
		ch.Name,
	)

	_, err = app.GetChannelByToken("test3_token")
	assert.NotNil(
		err,
	)
}
