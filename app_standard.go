package standard

import (
	api "github.com/james-nesbitt/coach-api"
)

// Application Simple implementation of the api.Application which maintains a list of Builder objects
type Application struct {
	builders api.Builders
}

// NewApplication Application constructor where you can pass in an optional Builders list
func NewApplication(builders api.Builders) *Application {
	if builders == nil {
		builders = NewBuilders().Builders()
	}
	return &Application{
		builders: builders,
	}
}

// AddBuilder Add a new builder to the application
func (sa *Application) AddBuilder(builder api.Builder) {
	sa.builders.Add(builder)
}

// Activate Enable some implementations on one of the already added Builder objects, with some settings
func (sa *Application) Activate(builderId string, implementations []string, settingsProvider api.SettingsProvider) error {
	if builder, err := sa.builders.Get(builderId); err == nil {
		return builder.Activate(implementations, settingsProvider)
	} else {
		return err
	}
}

// Validate Check that the Application is prepared to provide Operations
func (sa *Application) Validate() api.Result {
	res := NewResult()

	for _, bId := range sa.builders.Order() {
		builder, _ := sa.builders.Get(bId)
		res.Merge(builder.Validate())
	}

	return res.Result()
}

// Operations Provide an Operations list by collecting all of the activated Builder Operations
func (sa *Application) Operations() api.Operations {
	ops := NewOperations()

	for _, bId := range sa.builders.Order() {
		builder, _ := sa.builders.Get(bId)
		ops.Merge(builder.Operations())
	}

	return ops.Operations()
}
