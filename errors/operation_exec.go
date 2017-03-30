package errors

import (
	"context"
	"fmt"
)

type RequiredPropertyWasNotProvidedError struct {
	Key string
}

// Error return error string (interface: error)
func (rpwnpe RequiredPropertyWasNotProvidedError) Error() string {
	return fmt.Sprintf("REquired Property was not provided : %s", rpwnpe.Key)
}

type RequiredPropertyWasEmptyError struct {
	Key string
}

// Error return error string (interface: error)
func (rpwee RequiredPropertyWasEmptyError) Error() string {
	return fmt.Sprintf("REquired Property was not provided : %s", rpwee.Key)
}

type OperationTimedOut struct {
	Ctx context.Context
}

// Error return error string (interface: error)
func (oto OperationTimedOut) Error() string {
	return fmt.Sprintf("Required Property was not provided : %s", oto.Ctx.Err().Error())
}
