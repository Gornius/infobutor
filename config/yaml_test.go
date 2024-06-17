package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlParser_FromFile(t *testing.T) {
	yp := &YamlParser{}
	assert := assert.New(t)

	defaultConfig := GetDefaultConfig()
	bytes, err := yp.ToFile(defaultConfig)
	assert.Nil(err)

	configFromFile, err := yp.FromFile(bytes)
	assert.Nil(err)

	defaultConfigJson, err := json.Marshal(defaultConfig)
	assert.Nil(err)

	configFromFileJson, err := json.Marshal(configFromFile)
	assert.Nil(err)

	assert.Equal(
		defaultConfigJson,
		configFromFileJson,
	)
}

func TestYamlParser_ToFile(t *testing.T) {
	yp := &YamlParser{}
	assert := assert.New(t)

	defaultConfig := GetDefaultConfig()
	bytes, err := yp.ToFile(defaultConfig)
	assert.Nil(err)
	configFromFile, err := yp.FromFile(bytes)
	assert.Nil(err)

	defaultConfigJson, err := json.Marshal(defaultConfig)
	assert.Nil(err)

	configFromFileJson, err := json.Marshal(configFromFile)
	assert.Nil(err)

	assert.Equal(
		defaultConfigJson,
		configFromFileJson,
	)

}
