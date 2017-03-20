package standard

import (
	api "github.com/james-nesbitt/coach-api"
)

type Application struct {
	builders api.Builders
}

func (sa *Application) AddBuilder(builder api.Builder) {
	sa.builders.Add(builder)
}

func (sa *Application) Activate(builderId string, implementations []string, settingsProvider api.BuilderSettingsProvider) error {
	if builder, err := sa.builders.Get(builderId); err == nil {
		return builder.Activate(implementations, settingsProvider)
	} else {
		return err
	}
}

func (sa *Application) Validate() api.Result {
	res := NewResult()

	for _, bId := range sa.builders.Order() {
		builder, _ := sa.builders.Get(bId)
		res.Merge(builder.Validate())
	}

	return res.Result()
}

func (sa *Application) Operations() api.Operations {
	ops := NewOperations()

	for _, bId := range sa.builders.Order() {
		builder, _ := sa.builders.Get(bId)
		ops.Merge(builder.Operations())
	}

	return ops.Operations()
}
