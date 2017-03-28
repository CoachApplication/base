package null_test

import (
	"testing"

	null "github.com/CoachApplication/base/handler/null"
)

func TestBuilder_Operations(t *testing.T) {
	b := null.NewBuilder()

	b.Activate([]string{"config"}, nil)

	ops := b.Operations()

	if len(ops.Order()) == 0 {
		t.Error("NullBuilder didn't provide any Operations")
	}
}
