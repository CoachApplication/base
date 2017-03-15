package configyml

// YamlConfig a Config implementation that converts ConfigConnector read bytes 
// as YAML
type YamlConfig {
	cc handler_configconnector.ConfigConnector
}

// Get Applies config-connector reader data to an interface{}
func (yc *YamlConfig) Get(target interface{}) error {
	
}

// Set Uses a passed interface{} to crfeate yaml which is then saved to the ConfigConnector
func (yc *YamlConfig) Set(source interface{}) error {

}
