package config

import (
	"os"
	"testing"

	"github.com/gornius/infobutor/sender/manager"
	"github.com/stretchr/testify/assert"
)

func TestGetDefaultConfig(t *testing.T) {
	assert := assert.New(t)
	defaultConfig := GetDefaultConfig()

	assert.NotEmpty(defaultConfig)
}

func TestFromMap(t *testing.T) {
	assert := assert.New(t)

	configMap := GetDefaultConfig()
	defaultSinkToken := configMap["sinks"].(map[string]any)["default"].(map[string]any)["token"].(string)

	mgr, err := manager.New()
	assert.Nil(err)

	err = mgr.RegisterBuiltInSenders()
	assert.Nil(err)

	config, err := FromMap(mgr, configMap)
	assert.Nil(err)

	assert.Equal(
		defaultSinkToken,
		config.Sinks["default"].Token,
	)
}

func TestFromFile(t *testing.T) {
	assert := assert.New(t)

	defaultConfig := GetDefaultConfig()

	mgr, err := manager.New()
	mgr.RegisterBuiltInSenders()
	assert.Nil(err)

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
