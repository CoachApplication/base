package base_test

import (
	"testing"

	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
)

func Test_testPropertyId(t *testing.T) {
	prop := NewTestProperty("one", "One", "Prop One", "", base.OptionalPropertyUsage{}.Usage(), true)

	id := prop.Id()
	if id != "one" {
		t.Error("Incorrect Property Id() returned :", id)
	}
}

func Test_testPropertySet(t *testing.T) {
	prop := NewTestProperty("two", "Two", "Prop Two", "", base.OptionalPropertyUsage{}.Usage(), true)

	var val int = 2
	prop.Set(val)

	if getVal, ok := prop.Get().(int); !ok {
		t.Error("TestProperty returned incorrect val type")
	} else if getVal != 2 {
		t.Error("TestProperty returned incorrect val")
	}
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

	valid bool

	val interface{}
}

func NewTestProperty(id, label, description, help string, usage api.Usage, valid bool) *TestProperty {

	if usage == nil {
		usage = (&base.OptionalPropertyUsage{}).Usage()
	}

	return &TestProperty{
		id:          id,
		label:       label,
		description: description,
		help:        help,
		usage:       usage,
		valid:       valid,
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
func (tp *TestProperty) Validate() bool {
	return tp.valid
}
func (tp *TestProperty) Ui() api.Ui {
	return base.NewUi(
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
