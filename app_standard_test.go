package base_test

import (
	"testing"

	base "github.com/CoachApplication/coach-base"
)

func TestApplication_AddBuilder(t *testing.T) {
	app := base.NewApplication(nil)

	app.AddBuilder(NewTestBuilder("test.1").Builder())
	app.Activate("test.1", []string{"test"}, nil)

	ops := app.Operations()
	if len(ops.Order()) == 0 {
		t.Error("Application with activated Builder did not provide any Operations")
	}
}
