package property

import (
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
		return error(PropertyValWrongType{Id: rp.Id(), Type: rp.Type(), Val: val})
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
		return error(PropertyValWrongType{Id: wp.Id(), Type: wp.Type(), Val: val})
	}
}

// ReadWriterProperty Base Property for Properties that hold a io.ReadWriter
type ReadWriterProperty struct {
	val io.ReadWriter
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (rcp *ReadWriterProperty) Id() string {
	return "ReadWriterProperty"
}

// Type string label for content type of property value
func (rcp *ReadWriterProperty) Type() string {
	return "io.ReadWriter"
}

// Get retrieve a value from the Property
func (rcp *ReadWriterProperty) Get() interface{} {
	return interface{}(rcp.val)
}

// Set assign a value to the Property
func (rcp *ReadWriterProperty) Set(val interface{}) error {
	if typedVal, success := val.(io.ReadWriter); success {
		rcp.val = typedVal
		return nil
	} else {
		return error(PropertyValWrongType{Id: rcp.Id(), Type: rcp.Type(), Val: val})
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
		return error(PropertyValWrongType{Id: rcp.Id(), Type: rcp.Type(), Val: val})
	}
}
