package property

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

// Get retrieve a value from the Property
func (bp *BoolProperty) Get() (interface{}, error) {
	return interface{}(bp.val), nil
}

// Set assign a value to the Property
func (bp *BoolProperty) Set(val interface{}) error {
	if typedVal, success := val.(bool); success {
		bp.val = typedVal
		return nil
	} else {
		return error(PropertyValWrongType{Id: bp.Id(), Type: bp.Type(), Val: val})
	}
}
