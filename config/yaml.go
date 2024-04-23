package config

import (
	_ "embed"
	"os"

	"github.com/gornius/infobutor/sender/manager"
	"gopkg.in/yaml.v3"
)

type YamlParser struct {
}

func (yp *YamlParser) FromFile(senderManager *manager.Manager, file []byte) (map[string]any, error) {
	configMap := map[string]any{}

	if err := yaml.Unmarshal(file, configMap); err != nil {
		return nil, err
	}
	return configMap, nil
}

func (yp *YamlParser) ToFile(path string, configMap map[string]any) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return err
	}

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	if err := encoder.Encode(configMap); err != nil {
		return err
	}

	return nil
}
