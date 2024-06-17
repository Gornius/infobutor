package config

import (
	"os"
	"testing"

	"github.com/gornius/infobutor/sender/manager"
	"github.com/stretchr/testify/assert"
)

func TestDefaultLocation(t *testing.T) {
	assert := assert.New(t)
	loc := DefaultLocation()

	assert.NotEmpty(loc)
}

func TestGetDefaultConfig(t *testing.T) {
	assert := assert.New(t)
	defaultConfig := GetDefaultConfig()

	assert.NotEmpty(defaultConfig)
}

func TestFromMap(t *testing.T) {
	assert := assert.New(t)

	configMap := GetDefaultConfig()
	defaultChannelToken := configMap["channels"].(map[string]any)["default"].(map[string]any)["token"].(string)

	mgr := manager.NewWithAllBuiltIn()

	config, err := FromMap(mgr, configMap)
	assert.Nil(err)

	assert.Equal(
		defaultChannelToken,
		config.Channels["default"].Token,
	)
}

func TestFromFile(t *testing.T) {
	assert := assert.New(t)

	defaultConfig := GetDefaultConfig()

	mgr := manager.NewWithAllBuiltIn()
	parser := YamlParser{}

	tempFile, err := os.CreateTemp("", "gotests")
	if err != nil {
		assert.Fail("couldn't create temp dir")
	}

	bytes, err := parser.ToFile(defaultConfig)
	assert.Nil(err)
	tempFile.Write(*bytes)

	_, err = FromFile(mgr, tempFile.Name()+".yaml")

	assert.Nil(err)
}
