package property

import (
	"fmt"
)

type PropertyValWrongType struct {
	Id   string
	Type string
	Val  interface{}
}

func (pvwt PropertyValWrongType) Error() string {
	return fmt.Sprint(fmt.Sprintf("Wrong type of data was passed to %s Property set.  Expected %s", pvwt.Id, pvwt.Type), pvwt.Val)
}
