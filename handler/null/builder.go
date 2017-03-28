package null

import (
	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
	utils "github.com/CoachApplication/utils"
)

type Builder struct {
	imps utils.UniqueStringSlice
}

func NewBuilder() *Builder {
	return &Builder{
		imps: utils.UniqueStringSlice{},
	}
}

// Id provides a unique machine name for the Builder
func (nb *Builder) Id() string {
	return "null"
}

// SetParent Provides the API reference to the Builder which may use it's operations internally
func (nb *Builder) SetParent(api.API) {}

// Activate Enable keyed implementations, providing settings for those handler implementations
func (nb *Builder) Activate(imps []string, settings api.SettingsProvider) error {
	nb.imps.Merge(imps)
	return nil
}

// Validates Ask the builder if it is happy and willing to provide operations
func (nb *Builder) Validate() api.Result {
	return base.MakeSuccessfulResult()
}

// Operations provide any Builder user with a set of Operation objects
func (nb *Builder) Operations() api.Operations {
	ops := base.NewOperations()

	for _, imp := range nb.imps.Slice() {
		switch imp {
		case "config":
			ops.Merge(MakeConfigOperations(true, true))
		}
	}

	return ops.Operations()
}
