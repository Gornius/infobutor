package infobutor

import (
	"errors"
	"log"
	"net/http"

	"github.com/gornius/infobutor/sink"
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

type ReloadConfigRequest struct {
	Secret string `json:"secret"`
}

func NewApp() *App {
	app := new(App)
	app.Router = echo.New()

	app.Router.POST("/send/:sinkToken", func(c echo.Context) error {
		var msg message.Message
		sinkToken := c.Param("sinkToken")
		if err := c.Bind(&msg); err != nil {
			return c.NoContent(http.StatusBadRequest) // TODO: Implement error handling
		}

		sink, err := app.GetSinkByToken(sinkToken)
		if err != nil {
			return c.NoContent(http.StatusBadRequest) // TODO: Implement error handling
		}

		err = sink.Send(&msg)
		if err != nil {
			return err
		}

		return nil
	})

	app.Router.POST("/reload-config", func(c echo.Context) error {
		var body ReloadConfigRequest
		if err := c.Bind(&body); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if body.Secret != app.Config.Secret {
			reason := "bad secret given"
			return c.JSON(http.StatusBadRequest, ReloadConfigResponse{
				Success: false,
				Reason:  &reason,
			})
		}

		if err := app.ReloadConfig(); err != nil {
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
	app.LoadConfig(config.DefaultLocation())

	return app, nil
}

func (a *App) GetSinkByToken(token string) (*sink.Sink, error) {
	var sink *sink.Sink
	for _, ch := range a.Config.Sinks {
		if ch.Token == token {
			sink = ch
			break
		}
	}
	if sink == nil {
		return nil, errors.New("provided token does not exists on any of defined senders")
	}
	return sink, nil
}

func (a *App) LoadConfig(path string) error {
	a.configPath = path
	cfg, err := config.FromFile(a.SenderManager, a.configPath)
	if err != nil {
		return err
	}
	a.Config = cfg
	return nil
}

func (a *App) ReloadConfig() error {
	err := a.LoadConfig(a.configPath)
	if err != nil {
		return err
	}
	log.Println("Config reloaded!")
	return nil
}

func (a *App) GetConfigPath() string {
	return a.configPath
}
