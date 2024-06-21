package infobutor

import (
	"errors"
	"log"
	"net/http"

	"github.com/gornius/infobutor/channel"
	"github.com/gornius/infobutor/config"
	"github.com/gornius/infobutor/message"
	"github.com/gornius/infobutor/sender/manager"
	"github.com/labstack/echo/v4"
)

type App struct {
	Config        *config.Config
	configPath    string
	SenderManager *manager.Manager
	Router        *echo.Echo
}

type ReloadConfigResponse struct {
	Success bool    `json:"success"`
	Reason  *string `json:"reason"`
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

	app.Router.POST("/reload-config", func(c echo.Context) error {
		err := app.ReloadConfig()
		if err != nil {
			msg := err.Error()
			return c.JSON(http.StatusInternalServerError, &ReloadConfigResponse{
				Success: false,
				Reason:  &msg,
			})
		}

		return c.JSON(http.StatusOK, &ReloadConfigResponse{
			Success: true,
		})
	})

	return app
}

func NewDefaultApp() (*App, error) {
	app := NewApp()
	manager := manager.NewWithAllBuiltIn()
	app.SenderManager = manager
	app.SetConfigPath(config.DefaultLocation())

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

func (a *App) ReloadConfig() error {
	cfg, err := config.FromFile(a.SenderManager, a.configPath)
	if err != nil {
		return err
	}
	a.Config = cfg
	log.Println("Config reloaded!")
	return nil
}

func (a *App) SetConfigPath(configPath string) error {
	a.configPath = configPath

	err := a.ReloadConfig()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetConfigPath() string {
	return a.configPath
}
