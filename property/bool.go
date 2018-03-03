package property

import (
	base_errors "github.com/CoachApplication/base/errors"
)

// BoolProperty Base Property for Properties that hold boolean data
type BoolProperty struct {
	val bool
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (bp *BoolProperty) Id() string {
	return "BoolProperty"
}

// Type string label for content type of property value
func (bp *BoolProperty) Type() string {
	return "bool"
}

// Validate the Property
func (bp *BoolProperty) Validate() bool {
	return true
}

// Get retrieve a value from the Property
func (bp *BoolProperty) Get() interface{} {
	return interface{}(bp.val)
}

// Set assign a value to the Property
func (bp *BoolProperty) Set(val interface{}) error {
	if typedVal, success := val.(bool); success {
		bp.val = typedVal
		return nil
	} else {
		return error(base_errors.PropertyWrongValueTypeError{Id: bp.Id(), Type: bp.Type(), Val: val})
	}
}
