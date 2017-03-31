package errors

import "fmt"

type PropertyWrongValueTypeError struct {
	Id           string
	ExpectedType string
}

func (pwvte PropertyWrongValueTypeError) Error() string {
	return fmt.Sprintf("Property %s received the wrong value type.  Expected %s", pwvte.Id, pwvte.ExpectedType)
}
