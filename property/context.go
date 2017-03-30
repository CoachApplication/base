package property

import (
	"context"
	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
)

// StringProperty Base Property for Properties that hold a single string
type ContextProperty struct {
	val context.Context
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (cp *ContextProperty) Id() string {
	return "ContextProperty"
}

// Type string label for content type of property value
func (cp *ContextProperty) Type() string {
	return "context.Context"
}

// Type string label for content type of property value
func (cp *ContextProperty) Validate() api.Result {
	if cp.val == nil {
		return base.MakeFailedResult()
	} else {
		return base.MakeSuccessfulResult()
	}
}

// Get retrieve a value from the Property
func (cp *ContextProperty) Get() interface{} {
	return interface{}(cp.val)
}

// Set assign a value to the Property
func (cp *ContextProperty) Set(val interface{}) error {
	if typedVal, success := val.(context.Context); success {
		cp.val = typedVal
		return nil
	} else {
		return error(PropertyValWrongType{Id: cp.Id(), Type: cp.Type(), Val: val})
	}
}

const (
	PROPERTY_ID_CONTEXTLIMIT = "run.context"
)

type ContextLimitProperty struct {
	ContextProperty
}

func NewContextLimitProperty(ctx context.Context) *ContextLimitProperty {
	p := &ContextLimitProperty{}
	p.Set(ctx)
	return p
}

func (clp *ContextLimitProperty) Property() api.Property {
	return api.Property(clp)
}

func (clp *ContextLimitProperty) Id() string {
	return PROPERTY_ID_CONTEXTLIMIT
}

func (clp *ContextLimitProperty) Ui() api.Ui {
	return base.NewUi(
		clp.Id(),
		"Limit context",
		"Limit operation runtime based on a context",
		"",
	)
}

func (clp *ContextLimitProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}
