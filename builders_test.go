package base_test

import (
	"testing"

	base "github.com/CoachApplication/coach-base"
)

func TestSimpleBuilders_Add(t *testing.T) {
	builders := base.NewBuilders()

	builders.Add(NewTestBuilder("test"))

	bList := builders.Order()

	if len(bList) == 0 {
		t.Error("Builders did properly add a test Builder.  Builders shows no items added")
	} else if bList[0] != "test" {
		t.Error("Builders set Builder ID is not the same as the set Builder")
	}
}

func TestSimpleBuilders_Get(t *testing.T) {
	builders := base.NewBuilders()

	builders.Add(NewTestBuilder("test"))

	if getB, err := builders.Get("test"); err != nil {
		t.Error("Builders did not retrieve set Builder")
	} else if getB.Id() != "test" {
		t.Error("Builders retrieved Builder is not the same as the set Builder")
	}
}

func TestSimpleBuilders_Order(t *testing.T) {
	builders := base.NewBuilders()

	builders.Add(NewTestBuilder("test.1"))
	builders.Add(NewTestBuilder("test.2"))
	builders.Add(NewTestBuilder("test.3"))

	bList := builders.Order()

	if len(bList) == 0 {
		t.Error("Builders did properly add a test Builder.  Builders shows no items added")
	} else if bList[0] != "test.1" || bList[1] != "test.2" || bList[2] != "test.3" {
		t.Error("Builders set Builder ID is not the same as the set Builder")
	}
}
