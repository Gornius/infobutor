package main

import (
	"github.com/gornius/infobutor"
)

func main() {
	app, err := infobutor.NewDefaultApp()
	if err != nil {
		panic(err)
	}

	app.Router.Logger.Fatal(app.Router.Start(":3000"))
}
