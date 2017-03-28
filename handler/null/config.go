package null

import (
	api "github.com/CoachApplication/coach-api"
	base "github.com/CoachApplication/coach-base"
	"github.com/CoachApplication/coach-config"
)

// MakeConfigOperations make Null Operation for Config
func MakeConfigOperations(get, list bool) api.Operations {
	ops := base.NewOperations()

	if get {
		ops.Add(api.Operation(&GetOperation{}))
	}
	if list {
		ops.Add(api.Operation(&ListOperation{}))
	}

	return ops.Operations()
}

type GetOperation struct {
	config.GetOperationBase
	BaseOperation
}

type ListOperation struct {
	config.ListOperationBase
	BaseOperation
}
