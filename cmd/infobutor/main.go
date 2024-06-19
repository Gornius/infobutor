package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/gornius/infobutor"
	"github.com/spf13/cobra"
)

func main() {
	execute()
}

func execute() {
	rootCmd := &cobra.Command{
		Use:   "infobutor",
		Short: "Infobutor is a program that handles distributing messages to user-defined sources",
		Run: func(cmd *cobra.Command, args []string) {
			app, err := infobutor.NewDefaultApp()
			if err != nil {
				fmt.Println(err)
				return
			}
			app.Router.Logger.Fatal(app.Router.Start(":" + cmd.Flag("port").Value.String()))
		},
	}
	rootCmd.PersistentFlags().StringP("port", "p", "3000", "Port you want to run infobutor at")

	rootCmd.AddCommand(&cobra.Command{
		Use:   "reload",
		Short: "Reloads configuration",
		Run: func(cmd *cobra.Command, args []string) {
			port := cmd.Flag("port").Value.String()
			err := sendReloadRequest(port)
			if err != nil {
				fmt.Println(err)
			}
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func sendReloadRequest(port string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			msg := r.(*runtime.TypeAssertionError).Error()
			err = errors.New("Got bad response from server: " + msg)
		}
	}()

	data := map[string]any{}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	response, err := http.Post("http://127.0.0.1:"+port+"/reload-config", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	var responseData map[string]any
	json.NewDecoder(response.Body).Decode(&responseData)

	if !responseData["configReloaded"].(bool) {
		return errors.New("couldn't reload config")
	}

	return err
}
