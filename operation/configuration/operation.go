package configuration

import (
	coach_api "github.com/james-nesbitt/coach-api"
	coach_base "github.com/james-nesbitt/coach-base"
)

// GetOperation Operation that retrieves a single ScopedConfig for a config Key
type GetOperation struct{}

// Id Provide a unique machine name string
func (gon *GetOperation) Id() string {
	return "config.get"
}

// Usage Define UI metadata for the Operation
func (gon *GetOperation) Ui() coach_api.Ui {
	return coach_base.NewUi(
		gon.Id(),                                                                                     // Id
		"Get Configuration",                                                                          // Label
		"Retrieve scoped Configuration from a configuration backend",                                 // Description
		"Use this Operation to retrieve stored configuration from the system configuration backend.", // Help
	).Ui()
}

// Usage Define how the operations is intended to be used
func (gon *GetOperation) Usage() coach_api.Usage {
	return &(coach_base.InternalOperationUsage{}).Usage()
}

// ListOperation Operation that produces a list of Config keys
type ListOperation struct{}

// Id Provide a unique machine name string
func (lo *ListOperation) Id() string {
	return "config.get"
}

// Usage Define UI metadata for the Operation
func (lo *ListOperation) Ui() coach_api.Ui {
	return coach_base.NewUi(
		lo.Id(),                                                                                  // Id
		"List Configuration",                                                                     // Label
		"List Configurations avaialble from a configuration backend",                             // Description
		"Use this Operation to list stored configuration from the system configuration backend.", // Help
	).Ui()
}

// Usage Define how the operations is intended to be used
func (lo *ListOperation) Usage() coach_api.Usage {
	return &(coach_base.InternalOperationUsage{}).Usage()
}
