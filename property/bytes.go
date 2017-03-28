package property

import (
	api "github.com/CoachApplication/coach-api"
	base "github.com/CoachApplication/coach-base"
)

// ByteProperty Base Property for Properties that hold a single byte
type ByteProperty struct {
	val byte
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (bp *ByteProperty) Id() string {
	return "ByteProperty"
}

// Type string label for content type of property value
func (bp *ByteProperty) Type() string {
	return "byte"
}

// Get retrieve a value from the Property
func (bp *ByteProperty) Validate() api.Result {
	return base.MakeSuccessfulResult()
}

// Get retrieve a value from the Property
func (bp *ByteProperty) Get() interface{} {
	return interface{}(bp.val)
}

// Set assign a value to the Property
func (bp *ByteProperty) Set(val interface{}) error {
	if typedVal, success := val.(byte); success {
		bp.val = typedVal
		return nil
	} else {
		return error(PropertyValWrongType{Id: bp.Id(), Type: bp.Type(), Val: val})
	}
}

// ByteSliceProperty Base Property for Properties that hold a byte slice
type ByteSliceProperty struct {
	val []byte
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (bsp *ByteSliceProperty) Id() string {
	return "ByteSliceProperty"
}

// Type string label for content type of property value
func (bsp *ByteSliceProperty) Type() string {
	return "[]byte"
}

// Get retrieve a value from the Property
func (bsp *ByteSliceProperty) Validate() api.Result {
	return base.MakeSuccessfulResult()
}

// Get retrieve a value from the Property
func (bsp *ByteSliceProperty) Get() interface{} {
	return interface{}(bsp.val)
}

// Set assign a value to the Property
func (bsp *ByteSliceProperty) Set(val interface{}) error {
	if typedVal, success := val.([]byte); success {
		bsp.val = typedVal
		return nil
	} else {
		return error(PropertyValWrongType{Id: bsp.Id(), Type: bsp.Type(), Val: val})
	}
}

// StringProperty Base Property for Properties that hold a single string
type StringProperty struct {
	val string
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (sp *StringProperty) Id() string {
	return "StringProperty"
}

// Type string label for content type of property value
func (sp *StringProperty) Type() string {
	return "[]byte"
}

// Get retrieve a value from the Property
func (sp *StringProperty) Validate() api.Result {
	return base.MakeSuccessfulResult()
}

// Get retrieve a value from the Property
func (sp *StringProperty) Get() interface{} {
	return interface{}(sp.val)
}

// Set assign a value to the Property
func (sp *StringProperty) Set(val interface{}) error {
	if typedVal, success := val.(string); success {
		sp.val = typedVal
		return nil
	} else {
		return error(PropertyValWrongType{Id: sp.Id(), Type: sp.Type(), Val: val})
	}
}

// StringProperty Base Property for Properties that hold a single string
type StringSliceProperty struct {
	val []string
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (ssp *StringSliceProperty) Id() string {
	return "StringSliceProperty"
}

// Type string label for content type of property value
func (ssp *StringSliceProperty) Type() string {
	return "[]byte"
}

// Get retrieve a value from the Property
func (ssp *StringSliceProperty) Validate() api.Result {
	return base.MakeSuccessfulResult()
}

// Get retrieve a value from the Property
func (ssp *StringSliceProperty) Get() interface{} {
	return interface{}(ssp.val)
}

// Set assign a value to the Property
func (ssp *StringSliceProperty) Set(val interface{}) error {
	if typedVal, success := val.([]string); success {
		ssp.val = typedVal
		return nil
	} else {
		return error(PropertyValWrongType{Id: ssp.Id(), Type: ssp.Type(), Val: val})
	}
}
