package infobutor

import (
	"errors"
	"net/http"

	"github.com/gornius/infobutor/channel"
	"github.com/gornius/infobutor/config"
	"github.com/gornius/infobutor/message"
	"github.com/gornius/infobutor/sender/manager"
	"github.com/labstack/echo/v4"
)

type App struct {
	Config        *config.Config
	SenderManager *manager.Manager
	Router        *echo.Echo
}

func NewApp() *App {
	app := new(App)
	app.Router = echo.New()
	app.Router.POST("/send/:channelToken", func(c echo.Context) error {
		var msg message.Message
		channelToken := c.Param("channelToken")
		if err := c.Bind(&msg); err != nil {
			return c.NoContent(http.StatusBadRequest) // TODO: Implement error handling
		}

		channel, err := app.GetChannelByToken(channelToken)
		if err != nil {
			return c.NoContent(http.StatusBadRequest) // TODO: Implement error handling
		}

		err = channel.Send(&msg)
		if err != nil {
			return err
		}

		return nil
	})
	return app
}

func NewDefaultApp() (*App, error) {
	app := NewApp()
	manager := manager.NewWithAllBuiltIn()
	app.SenderManager = &manager

	config, err := config.FromFile(app.SenderManager, config.DefaultLocation())
	if err != nil {
		return nil, err
	}
	app.Config = config

	return app, nil
}

func (a *App) GetChannelByToken(token string) (*channel.Channel, error) {
	var channel *channel.Channel
	for _, ch := range a.Config.Channels {
		if ch.Token == token {
			channel = ch
			break
		}
	}
	if channel == nil {
		return nil, errors.New("provided token does not exists on any of defined senders")
	}
	return channel, nil
}
