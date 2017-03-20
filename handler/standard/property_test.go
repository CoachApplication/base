package standard

import (
	api "github.com/james-nesbitt/coach-api"
	"testing"
)

func Test_testPropertyId(t *testing.T) {

}

func Test_testPropertySet(t *testing.T) {

}

/**
 * Test structs
 */

// Base test struct to share Usage and Accessors
type testPropertyBase struct {
	val interface{}
}
func (tp *testPropertyBase) Usage() api.Usage {
	return OptionalPropertyUsage{}
}
func (tp *testPropertyBase) Get() interface{} {
	return tp.val
}
func (tp *testPropertyBase) Set(val interface{}) {
	tp.val = val
}

// First unique test Property
type testPropertyOne struct {
	testPropertyBase
}
func (tp *testPropertyOne) Id() string {
	return "test.1"
}
func (tp *testPropertyOne) Ui() api.Ui {
	return NewUi(
		tp.Id(),
		"test property",
		"This is a test property",
		"This test property can be used to test stuff",
	)
}

// Second unique test Property
type testPropertyTwo struct {
	testPropertyBase
}
func (tp *testPropertyTwo) Id() string {
	return "test.2"
}
func (tp *testPropertyTwo) Ui() api.Ui {
	return NewUi(
		tp.Id(),
		"test property",
		"This is a test property",
		"This test property can be used to test stuff",
	)
}

