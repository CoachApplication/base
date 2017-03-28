package null

import (
	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
	"github.com/CoachApplication/config"
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
