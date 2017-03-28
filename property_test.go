package standard

import (
	api "github.com/CoachApplication/coach-api"
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
type TestProperty struct {
	id          string
	label       string
	description string
	help        string

	usage api.Usage
	ui    api.Ui

	val interface{}
}

func NewTestProperty(id, label, description, help string, usage api.Usage) *TestProperty {

	if usage == nil {
		usage = (&OptionalPropertyUsage{}).Usage()
	}

	return &TestProperty{
		id:          id,
		label:       label,
		description: description,
		help:        help,
		usage:       usage,
	}
}

func (tp *TestProperty) Property() api.Property {
	return api.Property(tp)
}
func (tp *TestProperty) Id() string {
	return tp.id
}
func (tp *TestProperty) Type() string {
	return "interface{}"
}
func (tp *TestProperty) Usage() api.Usage {
	return tp.usage
}
func (tp *TestProperty) Validate() api.Result {
	res := NewResult()
	res.MarkSucceeded()
	res.MarkFinished()
	return res.Result()
}
func (tp *TestProperty) Ui() api.Ui {
	return NewUi(
		tp.id,
		tp.label,
		tp.description,
		tp.help,
	)
}
func (tp *TestProperty) Get() interface{} {
	return tp.val
}
func (tp *TestProperty) Set(val interface{}) error {
	tp.val = val
	return nil
}
