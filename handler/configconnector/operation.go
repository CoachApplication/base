package configconnector

// OperationBase a operation operation that holds a ConfigConector
type OperationBase struct {
	connector *ConfigConnector
}

func NewOperationBase(cc ConfigConnector) *OperationBase {
	return &OperationBase{
		connector: cc,
	}
}

// ConfigConnector retrieves a ConfigConnector
func (base OperationBase) ConfigConnector() ConfigConnector {
	return base.connector
}
