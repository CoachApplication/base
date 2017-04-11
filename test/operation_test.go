package test_test

import (
	"testing"

	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	"github.com/CoachApplication/base/test"
)

func TestTestOperation_Id(t *testing.T) {
	op := test.NewTestOperation("id", "label", "description", "help", nil, nil)

	if op.Id() != "id" {
		t.Error("TestOperation has the wrong Id() :", op.Id())
	}
}

func TestTestOperation_Ui(t *testing.T) {
	op := test.NewTestOperation("id", "label", "description", "help", nil, nil)
	ui := op.Ui()

	if ui.Id() != "id" {
		t.Error("TestOperation Ui gave the wrong Id(): ", ui.Id())
	} else if ui.Label() != "label" {
		t.Error("TestOperation Ui gave the wrong Label(): ", ui.Label())
	} else if ui.Description() != "description" {
		t.Error("TestOperation Ui gave the wrong Description(): ", ui.Description())
	} else if ui.Help() != "help" {
		t.Error("TestOperation Ui gave the wrong Help(): ", ui.Help())
	}
}

func TestTestOperation_Usage(t *testing.T) {
	op := test.NewTestOperation("id", "label", "description", "help", nil, nil)
	usage := op.Usage()

	if !usage.Allows(api.UsageOperationPublicView) {
		t.Error("TestOperation Usage default says that id doesn't allow view")
	} else if !usage.Allows(api.UsageOperationPublicExecute) {
		t.Error("TestOperation Usage default says that id doesn't allow execute")
	}

	op2 := test.NewTestOperation("id", "label", "description", "help", nil, base.InternalOperationUsage{}.Usage())
	usage2 := op2.Usage()

	if usage2.Allows(api.UsageOperationPublicView) {
		t.Error("TestOperation Usage internal says that id does allow view")
	} else if usage2.Allows(api.UsageOperationPublicExecute) {
		t.Error("TestOperation Usage internal says that id doesn allow execute")
	}
}

/**
 * @TODO decide how these tests should work
 */

func TestTestOperation_Validate(t *testing.T) {

}

func TestTestOperation_Exec(t *testing.T) {

}
