package config

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

type YamlParser struct {
}

func (yp *YamlParser) FromFile(bytes *[]byte) (map[string]any, error) {
	configMap := map[string]any{}

	if err := yaml.Unmarshal(*bytes, configMap); err != nil {
		return nil, err
	}
	return configMap, nil
}

func (yp *YamlParser) ToFile(configMap map[string]any) (*[]byte, error) {
	buf := new(bytes.Buffer)

	encoder := yaml.NewEncoder(buf)
	encoder.SetIndent(2)
	if err := encoder.Encode(configMap); err != nil {
		return nil, err
	}

	bytes := buf.Bytes()
	return &bytes, nil
}
