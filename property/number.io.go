package property

import (
	base_errors "github.com/CoachApplication/base/errors"
)

// IntProperty Base Property for Properties that hold an integer
type IntProperty struct {
	val int
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (ip *IntProperty) Id() string {
	return "IntProperty"
}

// Type string label for content type of property value
func (ip *IntProperty) Type() string {
	return "int"
}

// Get retrieve a value from the Property
func (ip *IntProperty) Get() interface{} {
	return interface{}(ip.val)
}

// Set assign a value to the Property
func (ip *IntProperty) Set(val interface{}) error {
	if typedVal, success := val.(int); success {
		ip.val = typedVal
		return nil
	} else {
		return error(base_errors.PropertyWrongValueTypeError{Id: ip.Id(), Type: ip.Type(), Val: val})
	}
}

// FloatProperty Base Property for Properties that hold an integer
type FloatProperty struct {
	val float64
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (fp *FloatProperty) Id() string {
	return "FloatProperty"
}

// Type string label for content type of property value
func (fp *FloatProperty) Type() string {
	return "float64"
}

// Get retrieve a value from the Property
func (fp *FloatProperty) Get() interface{} {
	return interface{}(fp.val)
}

// Set assign a value to the Property
func (fp *FloatProperty) Set(val interface{}) error {
	if typedVal, success := val.(float64); success {
		fp.val = typedVal
		return nil
	} else {
		return error(base_errors.PropertyWrongValueTypeError{Id: fp.Id(), Type: fp.Type(), Val: val})
	}
}
