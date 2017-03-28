package base_test

import (
	api "github.com/CoachApplication/coach-api"
	base "github.com/CoachApplication/coach-base"
)

// TestBuilder that can be used for other testing
type TestBuilder struct {
	id              string
	parent          api.API
	implementations []string
}

// NewTestBuilder Constructor for TestBuilder
func NewTestBuilder(id string) *TestBuilder {
	return &TestBuilder{
		id: id,
	}
}

// Builder Explicitly convert this to an api.Builder
func (tb *TestBuilder) Builder() api.Builder {
	return api.Builder(tb)
}

// Id provides a unique machine name for the Builder
func (tb *TestBuilder) Id() string {
	return tb.id
}

// SetParent Provides the API reference to the Builder which may use it's operations internally
func (tb *TestBuilder) SetParent(API api.API) {
	tb.parent = API
}

// Activate Enable keyed implementations, providing settings for those handler implementations
func (tb *TestBuilder) Activate(implementations []string, settings api.SettingsProvider) error {
	tb.implementations = append(tb.implementations, implementations...)
	return nil
}

// Validates Ask the builder if it is happy and willing to provide operations
func (tb *TestBuilder) Validate() api.Result {
	res := base.NewResult()
	res.MarkSucceeded()
	res.MarkFinished()
	return res.Result()
}

// Operations provide any Builder user with a set of Operation objects
func (tb *TestBuilder) Operations() api.Operations {
	ops := base.NewOperations()

	for _, imp := range tb.implementations {
		ops.Add(NewTestOperation(imp+".1", imp+"Test Operation 1", "", "", nil, nil).Operation())
	}

	return ops.Operations()
}
