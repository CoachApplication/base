package base_test

import (
	"github.com/CoachApplication/base"
	"github.com/CoachApplication/base/test"
	"testing"
)

func TestProperties_AddandGet(t *testing.T) {
	props := base.NewProperties()
	id := "test.test"
	props.Add(test.NewTestProperty(id, "Test Property", "", "", nil, true).Property())

	if getProp, err := props.Get(id); err != nil {
		t.Error("Properties gave an error when retrieving a valid property id: ", id, err.Error())
	} else if getProp.Id() != id {
		t.Error("Properties returned the wrong property: ", id, getProp)
	}
}

func TestProperties_Order(t *testing.T) {
	props := base.NewProperties()
	props.Add(test.NewTestProperty("one", "Test Property one", "", "", nil, true).Property())
	props.Add(test.NewTestProperty("two", "Test Property one", "", "", nil, true).Property())
	props.Add(test.NewTestProperty("three", "Test Property one", "", "", nil, true).Property())
	props.Add(test.NewTestProperty("four", "Test Property one", "", "", nil, true).Property())

	if list := props.Order(); len(list) == 0 {
		t.Error("Properties returned an empty list")
	} else if len(list) != 4 {
		t.Error("Properties returned the wrong number of ids: ", list)
	} else if !(list[0] == "one" && list[1] == "two" && list[2] == "three" && list[3] == "four") {
		t.Error("Properties returned the wrong ids in the wrong order: ", list)
	}
}
