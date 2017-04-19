package property

import (
	base_errors "github.com/CoachApplication/base/errors"
	"io"
)

// ReaderProperty Base Property for Properties that hold a io.Reader
type ReaderProperty struct {
	val io.Reader
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (rp *ReaderProperty) Id() string {
	return "ReaderProperty"
}

// Type string label for content type of property value
func (rp *ReaderProperty) Type() string {
	return "io.Reader"
}

// Validate boolean for if the Property value is valid (not nil)
func (rp *ReaderProperty) Validate() bool {
	return rp.val != nil
}

// Get retrieve a value from the Property
func (rp *ReaderProperty) Get() interface{} {
	return interface{}(rp.val)
}

// Set assign a value to the Property
func (rp *ReaderProperty) Set(val interface{}) error {
	if typedVal, success := val.(io.Reader); success {
		rp.val = typedVal
		return nil
	} else {
		return error(base_errors.PropertyWrongValueTypeError{Id: rp.Id(), Type: rp.Type(), Val: val})
	}
}

// WriterProperty Base Property for Properties that hold a io.Reader
type WriterProperty struct {
	val io.Reader
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (wp *WriterProperty) Id() string {
	return "WriterProperty"
}

// Type string label for content type of property value
func (wp *WriterProperty) Type() string {
	return "io.Reader"
}

// Validate boolean for if the Property value is valid (not nil)
func (wp *WriterProperty) Validate() bool {
	return wp.val != nil
}

// Get retrieve a value from the Property
func (wp *WriterProperty) Get() interface{} {
	return interface{}(wp.val)
}

// Set assign a value to the Property
func (wp *WriterProperty) Set(val interface{}) error {
	if typedVal, success := val.(io.Reader); success {
		wp.val = typedVal
		return nil
	} else {
		return error(base_errors.PropertyWrongValueTypeError{Id: wp.Id(), Type: wp.Type(), Val: val})
	}
}

// ReadWriterProperty Base Property for Properties that hold a io.ReadWriter
type ReadWriterProperty struct {
	val io.ReadWriter
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (rwp *ReadWriterProperty) Id() string {
	return "ReadWriterProperty"
}

// Type string label for content type of property value
func (rwp *ReadWriterProperty) Type() string {
	return "io.ReadWriter"
}

// Validate boolean for if the Property value is valid (not nil)
func (rwp *ReadWriterProperty) Validate() bool {
	return rwp.val != nil
}

// Get retrieve a value from the Property
func (rwp *ReadWriterProperty) Get() interface{} {
	return interface{}(rwp.val)
}

// Set assign a value to the Property
func (rwp *ReadWriterProperty) Set(val interface{}) error {
	if typedVal, success := val.(io.ReadWriter); success {
		rwp.val = typedVal
		return nil
	} else {
		return error(base_errors.PropertyWrongValueTypeError{Id: rwp.Id(), Type: rwp.Type(), Val: val})
	}
}

// ReadCloserProperty Base Property for Properties that hold a io.ReadCloser
type ReadCloserProperty struct {
	val io.ReadCloser
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (rcp *ReadCloserProperty) Id() string {
	return "ReadCloserProperty"
}

// Type string label for content type of property value
func (rcp *ReadCloserProperty) Type() string {
	return "io.ReadWriter"
}

// Validate boolean for if the Property value is valid (not nil)
func (rcp *ReadCloserProperty) Validate() bool {
	return rcp.val != nil
}

// Get retrieve a value from the Property
func (rcp *ReadCloserProperty) Get() interface{} {
	return interface{}(rcp.val)
}

// Set assign a value to the Property
func (rcp *ReadCloserProperty) Set(val interface{}) error {
	if typedVal, success := val.(io.ReadCloser); success {
		rcp.val = typedVal
		return nil
	} else {
		return error(base_errors.PropertyWrongValueTypeError{Id: rcp.Id(), Type: rcp.Type(), Val: val})
	}
}
