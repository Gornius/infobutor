package infobutor

import (
	"os"
	"testing"

	"github.com/gornius/infobutor/config"
	"github.com/gornius/infobutor/sender"
	"github.com/gornius/infobutor/sink"
	"github.com/stretchr/testify/assert"
)

func TestGetSinkByToken(t *testing.T) {
	assert := assert.New(t)
	var err error

	configFile, err := writeDefaultConfigToTemp(config.GetDefaultConfig())
	if err != nil {
		assert.Fail("couldn't write a temp config file")
	}

	app, err := createAppForTesting(configFile)
	app.Config = &config.Config{
		Sinks: map[string]*sink.Sink{
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

	var ch *sink.Sink

	ch, err = app.GetSinkByToken("test1_token")
	assert.Nil(
		err,
	)
	assert.Equal(
		"test1_name",
		ch.Name,
	)

	ch, err = app.GetSinkByToken("test2_token")
	assert.Nil(
		err,
	)
	assert.Equal(
		"test2_name",
		ch.Name,
	)

	_, err = app.GetSinkByToken("test3_token")
	assert.NotNil(
		err,
	)
}

func TestLoadConfig(t *testing.T) {
	assert := assert.New(t)

	configFile, err := writeDefaultConfigToTemp(config.GetDefaultConfig())
	if err != nil {
		assert.Fail("couldn't write a temp config file")
	}
	app, err := createAppForTesting(configFile)

	assert.Nil(err)

	assert.Equal(
		configFile,
		app.configPath,
	)
}

func TestReloadConfig(t *testing.T) {
	assert := assert.New(t)

	cfg := config.GetDefaultConfig()
	cfg["secret"] = "test1"
	configFile, err := writeDefaultConfigToTemp(cfg)
	if err != nil {
		assert.Fail("couldn't write a temp config file")
	}

	cfg["secret"] = "test2"
	configFile2, err := writeDefaultConfigToTemp(cfg)
	if err != nil {
		assert.Fail("couldn't write a temp config file")
	}

	app, err := createAppForTesting(configFile)
	if err != nil {
		assert.Fail("couldn't create app instance")
	}

	assert.Equal(
		"test1",
		app.Config.Secret,
	)

	app.configPath = configFile2
	err = app.ReloadConfig()
	assert.Nil(err)

	assert.Equal(
		"test2",
		app.Config.Secret,
	)
}

func writeDefaultConfigToTemp(configMap map[string]any) (string, error) {
	tempFile, err := os.CreateTemp("", "gotests*.yaml")
	if err != nil {
		return "", err
	}

	parser := config.YamlParser{}
	bytes, err := parser.ToFile(configMap)
	if err != nil {
		return "", err
	}
	tempFile.Write(*bytes)

	return tempFile.Name(), nil
}

func createAppForTesting(configFile string) (*App, error) {
	app, err := NewApp(
		WithBuiltInSenders(),
		WithConfigPath(configFile),
	)

	if err != nil {
		return nil, err
	}

	return app, nil
}
