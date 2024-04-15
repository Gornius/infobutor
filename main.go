package main

import (
	"fmt"

	"github.com/gornius/infobutor/config"
	"github.com/gornius/infobutor/message"
	"github.com/gornius/infobutor/sender/manager"
)

func main() {
	senderManager := manager.WithAllBuiltIn()
	config, err := config.FromFile(&senderManager, config.DefaultLocation())
	if err != nil {
		fmt.Println("Failed to load config")
		return
	}

	message := message.Message{
		Title:   "TEST_TITLE",
		Content: `Test z konfiga :) 🥳🥳🥳🥳🥳`,
		Origin:  "TEST_ORIGIN",
	}

	for _, ch := range config.Channels {
		ch.Send(&message)
	}
}
