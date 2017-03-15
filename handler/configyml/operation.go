package configyml

// ConfigConnectorOperations return three operations from a config connect, the Get, Set and List operations
func YmlConfigConnectorOperations(cc *ConfigConnector) (api_operation.Operations) {
	ops := api_operation.SimpleOperations{}

	base := handler_configconnector.OperationBase{connector: cc}

	ops.add(api_operation.Operation(&GetOperation{OperationBase: *base}))
	ops.add(api_operation.Operation(&ListOperation{OperationBase: *base}))

	return api_operation.Operations(ops)
}


// ListOperation list all configs for a config-connector
type ListOperation struct {
	api_base_config.ListOperation
	handler_configconnector.OperationBase
}

// Exec runs the List operation
func (lo *ListOperation) Exec(props api_property.Properties) api_result.Result {
	res := api_result.SimpleResult{}



	return api_result.Result(res)	
}


// GetOperation Retrieve a scoped-config for config-connector
type GetOperation struct {
	api_base_config.GetOperation
	handler_configconnector.OperationBase
}

// Exec runs the Get operation
func (go *GetOperation) Exec(props api_property.Properties) api_result.Result {
	res := api_result.SimpleResult{}

	return api_result.Result(res)
}
