package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gornius/infobutor/channel"
	"github.com/gornius/infobutor/config"
	"github.com/gornius/infobutor/message"
	"github.com/gornius/infobutor/sender/manager"
	"github.com/labstack/echo/v4"
)

func main() {
	senderManager := manager.NewWithAllBuiltIn()
	config, err := config.FromFile(&senderManager, config.DefaultLocation())
	if err != nil {
		fmt.Println("Failed to load config")
		return
	}

	router := echo.New()
	router.POST("/send/:channelToken", func(c echo.Context) error {
		var msg message.Message
		channelToken := c.Param("channelToken")
		if err := c.Bind(&msg); err != nil {
			return c.NoContent(http.StatusBadRequest) // TODO: Implement error handling
		}

		var channel *channel.Channel
		for _, ch := range config.Channels {
			if ch.Token == channelToken {
				channel = ch
				break
			}
		}

		if channel == nil {
			return c.NoContent(http.StatusBadRequest) // TODO: Implement error handling
		}

		wg := sync.WaitGroup{}
		for _, channel := range channel.Senders {
			wg.Add(1)
			go func() {
				defer wg.Done()
				channel.Send(msg)
			}()
		}
		wg.Wait()

		return nil
	})

	router.Logger.Fatal(router.Start(":3000"))
}
