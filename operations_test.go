package base_test

import (
	"testing"

	base "github.com/CoachApplication/base"
)

func TestOperations_Add(t *testing.T) {
	ops := base.NewOperations()

	ops.Add(NewTestOperation("test.1", "", "", "", nil, nil).Operation())

	list := ops.Order()
	if len(list) == 0 {
		t.Error("Test Operations did not hold any operations after adding one")
	} else if get, err := ops.Get("test.1"); err != nil {
		t.Error("Test Operations was not able to retrieve an added operation", err.Error())
	} else if get.Id() != "test.1" {
		t.Error("Test Operations did not retreive that added operation, wrong id given", get.Id())
	}
}

func TestOperations_Get(t *testing.T) {
	ops := base.NewOperations()

	ops.Add(NewTestOperation("test.1", "", "", "", nil, nil).Operation())

	list := ops.Order()
	if len(list) == 0 {

	} else if get, err := ops.Get("test.1"); err != nil {
		t.Error("Test Operations was not able to retrieve an added operation", err.Error())
	} else if get.Id() != "test.1" {
		t.Error("Test Operations did not retreive that added operation, wrong id given", get.Id())
	}
}

func TestOperations_List(t *testing.T) {
	ops := base.NewOperations()

	ops.Add(NewTestOperation("test.1", "", "", "", nil, nil).Operation())
	ops.Add(NewTestOperation("test.2", "", "", "", nil, nil).Operation())
	ops.Add(NewTestOperation("test.3", "", "", "", nil, nil).Operation())

	list := ops.Order()
	if len(list) == 0 {
		t.Error("Test Operations did not hold any operations after some were added")
	} else if len(list) != 3 {
		t.Error("Test Operations did not hold the correct number of Operation")
	} else if list[0] != "test.1" || list[1] != "test.2" || list[2] != "test.3" {
		t.Error("Test Operations did not return the proper ordered ids")
	}
}
