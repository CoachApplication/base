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

func NewSuccessfulValidOperation(id string) *TestOperation {
	success := &SuccessfulOperationProperty{}
	valid:= &ValidOperationProperty{}

	success.Set(true)
	valid.Set(true)

	props := base.NewProperties()
	props.Add(success.Property())
	props.Add(valid.Property())

	return NewTestOperation(id, "TestSuccess", "", "", props.Properties(), nil)
}

func NewFailedValidOperation(id string) *TestOperation {
	success := &SuccessfulOperationProperty{}
	valid:= &ValidOperationProperty{}

	success.Set(false)
	valid.Set(true)

	props := base.NewProperties()
	props.Add(success.Property())
	props.Add(valid.Property())

	return NewTestOperation(id, "TestFailed", "", "", props.Properties(), nil)
}

func NewInValidOperation(id string) *TestOperation {
	success := &SuccessfulOperationProperty{}
	valid:= &ValidOperationProperty{}

	success.Set(false)
	valid.Set(false)

	props := base.NewProperties()
	props.Add(success.Property())
	props.Add(valid.Property())

	return NewTestOperation(id, "TestInvalud", "", "", props.Properties(), nil)
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
	success := &SuccessfulOperationProperty{}
	valid:= &ValidOperationProperty{}

	success.Set(false)
	valid.Set(false)

	props := base.NewProperties()
	props.Add(success.Property())
	props.Add(valid.Property())

	props.Merge(to.props)

	return props.Properties()
}

// Validate Validate that the Operation can Execute if passed proper Property data
func (to *TestOperation) Validate(props api.Properties) api.Result {
	res := base.NewResult()

	go func (props api.Properties) {
		if validProp, err := props.Get(PROPERTY_ID_OPERATIONVALID); err != nil {
			res.MarkFailed()
			res.AddError(err)
		} else if valid, good := validProp.Get().(bool); !good {
			res.MarkFailed()
		} else if !valid {
			res.AddProperty(validProp)
			res.MarkFailed()
		} else {
			res.AddProperty(validProp)
			res.MarkSucceeded()
		}
		res.MarkFinished()
	}(props)

	return res.Result()
}

//Exec runs the operation from a Properties set, and return a result
func (to *TestOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func (props api.Properties) {
		if successProp, err := props.Get(PROPERTY_ID_OPERATIONSUCCESS); err != nil {
			res.MarkFailed()
			res.AddError(err)
		} else if success, good := successProp.Get().(bool); !good {
			res.MarkFailed()
		} else if !success {
			res.AddProperty(successProp)
			res.MarkFailed()
		} else {
			res.AddProperty(successProp)
			res.MarkSucceeded()
		}
		res.MarkFinished()
	}(props)

	return res.Result()
}
