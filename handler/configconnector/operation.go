package configconnector

// OperationBase a base operation that holds a ConfigConector
type OperationBase {
	connector *ConfigConnector
}

// ConfigConnector retrieves a ConfigConnector
func (base OperationBase) ConfigConnector() ConfigConnector {
	return base.connector
}
