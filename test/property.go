package test

import (
	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
	"github.com/CoachApplication/base/property"
)

const (
	PROPERTY_ID_OPERATIONVALID = "test.validop"
)

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

type ValidOperationProperty struct {
	property.BoolProperty
}
func (vop *ValidOperationProperty) Property() api.Property {
	return api.Property(vop)
}
func (vop *ValidOperationProperty) Id() string {
	return PROPERTY_ID_OPERATIONVALID
}
func (vop *ValidOperationProperty) Usage() api.Usage {
	return base.RequiredPropertyUsage{}.Usage()
}
func (vop *ValidOperationProperty) Validate() bool {
	return true
}
func (vop *ValidOperationProperty) Ui() api.Ui {
	return base.NewUi(
		vop.Id(),
		"Valid",
		"Test Operation is valid",
		"",
	)
}
