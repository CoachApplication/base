package standard

import (
	"testing"
)

func TestApplication_AddBuilder(t *testing.T) {
	app := NewApplication(nil)

	app.AddBuilder(NewTestBuilder("test.1").Builder())
	app.Activate("test.1", []string{"test"}, nil)

	ops := app.Operations()
	if len(ops.Order()) == 0 {
		t.Error("Application with activated Builder did not provide any Operations")
	}
}

//
//func TestSecuredApplication_Operations(t *testing.T) {
//	authOp := NewAuthOperation(true)
//	app := NewSecuredApplication(authOp, "test.authorization.operation", "test.authorization.success")
//
//	app.AddBuilder(NewTestBuilder("test.1").Builder())
//	app.Activate("test.1", []string{"test"}, nil)
//
//	ops := app.Operations()
//	if len(ops.Order()) == 0 {
//		t.Error("Application with activated Builder did not provide any Operations")
//	}
//
//	if op, err := ops.Get(ops.Order()[0]); err != nil {
//		t.Error("SecuredApplication did not return the correct Operation", err.Error())
//	} else {
//		res := op.Exec(op.Properties())
//		<-res.Finished()
//
//		res.Success()
//
//		resProps := res.Properties()
//		if successProp, err := resProps.Get("test.authorization.success"); err != nil {
//			t.Error("Secured application did not provide any authorization success property", resProps)
//		} else if !successProp.Get().(bool) {
//			t.Error("Secured application was supposed to authorize test Operation, but it didn't")
//		}
//
//	}
//}
