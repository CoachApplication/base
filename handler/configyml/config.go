package configyml

import (
	base_config "github.com/james-nesbitt/coach-base/operation/configuration"
)

// YamlConfig a Config implementation that converts ConfigConnector read bytes
// as YAML
type YamlConfig struct {
	cc base_config.Connector
}

// Get Applies config-connector reader data to an interface{}
func (yc *YamlConfig) Get(target interface{}) error {
	return nil
}

// Set Uses a passed interface{} to create yaml which is then saved to the ConfigConnector
func (yc *YamlConfig) Set(source interface{}) error {
	return nil
}
