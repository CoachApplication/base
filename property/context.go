package property

import "context"

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
