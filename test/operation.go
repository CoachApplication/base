package test

import (
	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
)

// TestOperation used for testing
type TestOperation struct {
	id          string
	label       string
	description string
	help        string

	props api.Properties
	usage api.Usage
}

// Constructor for TestOperation
func NewTestOperation(id, label, description, help string, props api.Properties, usage api.Usage) *TestOperation {
	if id == "" {
		id = "test"
	}
	if props == nil {
		props = base.NewProperties().Properties()
	}
	if usage == nil {
		usage = (&base.ExternalOperationUsage{}).Usage()
	}

	return &TestOperation{
		id:          id,
		label:       label,
		description: description,
		help:        help,
		props:       props,
		usage:       usage,
	}
}

// Id Provide a unique machine name string identifier
func (to *TestOperation) Operation() api.Operation {
	return api.Operation(to)
}

// Id Provide a unique machine name string identifier
func (to *TestOperation) Id() string {
	return to.id
}

// UI Return a UI interaction definition for the Operation
func (to *TestOperation) Ui() api.Ui {
	return base.NewUi(
		to.id,
		to.label,
		to.description,
		to.help,
	).Ui()
}

// Usage Define how the Operation is intended to be executed.
func (to *TestOperation) Usage() api.Usage {
	return to.usage
}

// Properties provide the expected Operation with default values
func (to *TestOperation) Properties() api.Properties {
	return to.props
}

// Validate Validate that the Operation can Execute if passed proper Property data
func (to *TestOperation) Validate(props api.Properties) api.Result {
	return base.MakeSuccessfulResult()
}

//Exec runs the operation from a Properties set, and return a result
func (to *TestOperation) Exec(props api.Properties) api.Result {
	return base.MakeSuccessfulResult()
}

