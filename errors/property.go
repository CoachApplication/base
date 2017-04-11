package errors

import "fmt"

type PropertyWrongValueTypeError struct {
	Id   string
	Type string
	Val  interface{}
}

func (pwvte PropertyWrongValueTypeError) Error() string {
	return fmt.Sprint(fmt.Sprintf("Property %s received the wrong value type.  Expected %s [%s]", pwvte.Id, pwvte.Type), pwvte.Val)
}
