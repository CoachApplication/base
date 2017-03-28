package null_test

import (
	"testing"

	"github.com/CoachApplication/coach-base"
	null "github.com/CoachApplication/coach-base/handler/null"
	config "github.com/CoachApplication/coach-config"
)

func TestMakeConfigOperations(t *testing.T) {
	ops := null.MakeConfigOperations(true, true)

	if _, err := ops.Get(config.OPERATION_KEY_CONFIG_GET); err != nil {
		t.Error("Null An error occured retreiving GetOperation: ", err.Error())
	}
	if _, err := ops.Get(config.OPERATION_KEY_CONFIG_LIST); err != nil {
		t.Error("Null An error occured retreiving ListOperation: ", err.Error())
	}
}

func TestGetOperation_Exec(t *testing.T) {
	ops := null.MakeConfigOperations(true, true)

	if lOp, err := ops.Get(config.OPERATION_KEY_CONFIG_GET); err != nil {
		t.Error("Null An error occured retreiving GetOperation: ", err.Error())
	} else {
		res := lOp.Exec(base.NewProperties().Properties())
		<-res.Finished()

		if res.Success() {
			t.Error("Null GetOperation returned a successful result")
		}
	}
}

func TestListOperation_Exec(t *testing.T) {
	ops := null.MakeConfigOperations(true, true)

	if lOp, err := ops.Get(config.OPERATION_KEY_CONFIG_LIST); err != nil {
		t.Error("Null An error occured retreiving ListOperation: ", err.Error())
	} else {
		res := lOp.Exec(base.NewProperties().Properties())
		<-res.Finished()

		if res.Success() {
			t.Error("Null ListOperation returned a successful result")
		}
	}
}
