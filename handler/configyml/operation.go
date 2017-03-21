package configyml

import (
	coach_api "github.com/james-nesbitt/coach-api"
	coach_base "github.com/james-nesbitt/coach-base"
	coach_base_config "github.com/james-nesbitt/coach-base/operation/configuration"
	handler_configconnector "github.com/james-nesbitt/coach-base/handler/configconnector"
)

// ConfigConnectorOperations return three operations from a config connect, the Get, Set and List operations
func YmlConfigConnectorOperations(cc *handler_configconnector.ConfigConnector) coach_api.Operations {
	ops := coach_base.NewOperations()

	base := handler_configconnector.NewOperationBase(cc)

	ops.Add(coach_api.Operation(&GetOperation{OperationBase: *base}))
	ops.Add(coach_api.Operation(&ListOperation{OperationBase: *base}))

	return ops.Operations()
}

// ListOperation list all configs for a config-connector
type ListOperation struct {
	coach_base_config.ListOperation
	handler_configconnector.OperationBase
}

// Exec runs the List operation
func (lo *ListOperation) Exec(props coach_api.Properties) coach_api.Result {
	res := coach_base.NewResult()

	return res.Result()
}

// GetOperation Retrieve a scoped-config for config-connector
type GetOperation struct {
	coach_base_config.GetOperation
	handler_configconnector.OperationBase
}

// Exec runs the Get operation
func (gon *GetOperation) Exec(props coach_api.Properties) coach_api.Result {
	res := coach_base.NewResult()

	return res.Result()
}
