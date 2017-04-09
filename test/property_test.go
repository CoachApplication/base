package test_test

import (
	"testing"
	"github.com/CoachApplication/base"
	"github.com/CoachApplication/base/test"
)

func TestTestProperty_Id(t *testing.T) {
	prop := test.NewTestProperty("one", "One", "Prop One", "", base.OptionalPropertyUsage{}.Usage(), true)

	id := prop.Id()
	if id == "" {
		t.Error("Property Id() is empty")
	} else if id != "one" {
		t.Error("Incorrect Property Id() returned :", id)
	}
}

func TestTestProperty_SetGet(t *testing.T) {
	prop := test.NewTestProperty("two", "Two", "Prop Two", "", base.OptionalPropertyUsage{}.Usage(), true)

	var val int = 2
	prop.Set(val)

	if getVal, ok := prop.Get().(int); !ok {
		t.Error("TestProperty returned incorrect val type")
	} else if getVal != 2 {
		t.Error("TestProperty returned incorrect val")
	}
}

func TestTestProperty_Validate(t *testing.T) {
	prop := test.NewTestProperty("two", "Two", "Prop Two", "", base.OptionalPropertyUsage{}.Usage(), true)
	if !prop.Validate() {
		t.Error("TestProperty thinks it is invalid when it is valid")
	}

	prop = test.NewTestProperty("two", "Two", "Prop Two", "", base.OptionalPropertyUsage{}.Usage(), false)
	if prop.Validate() {
		t.Error("TestProperty thinks it is valid when it is invalid")
	}
}