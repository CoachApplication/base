package standard

import (
	"errors"
	"testing"

	api "github.com/james-nesbitt/coach-api"
)

// TestNewResult Test that the NewResult() function returns a struct that implements the api.Result interface
func TestNewResult(t *testing.T) {
	res := NewResult()

	var i interface{} = res
	if _, isResult:= i.(api.Result); !isResult {
		t.Error("NewResult did not provide a struct which is a valid api.Result implementation", res)
	}
}

// TestResult_AddProperty test adding an retrieving properties to a Result
func TestResult_AddProperty(t *testing.T) {
	res := NewResult()

	newProp := api.Property(&testPropertyOne{})
	newProp.Set(interface{}("one")) // we don't test the set/get here, do that in the property_test.go

	res.AddProperty(newProp)

	newProp.Set(interface{}("two"))

	if getProp, found := res.Properties().Get(newProp.Id()); found == nil {
		if val, isString := getProp.Get().(string); isString {
			switch val {
			case "one":
				t.Error("Result retrieve property has not maintained its reference", newProp.Id())
			case "two":
				// Success
			default:
				t.Error("Result did not retrieve the correct property that we added to it", newProp.Id())
			}
		} else {
			t.Error("Result did not retrieve the correct property that we added to it", newProp.Id())
		}
	} else {
		t.Error("Result did not retrieve a property that we added to it", newProp.Id())
	}
}

// TestResult_AddError Test that if we add an error, we get it back
func TestResult_AddError(t *testing.T) {
	res := NewResult()

	newErr := errors.New("test")
	res.AddError(newErr)

	errs := res.Errors()
	if len(errs) == 0 {
		t.Error("Result object did not return any errors, even though we added one to it", errs)
	} else if errs[0].Error() != "test" {
		t.Error("Result object did not return the error that we added to it", errs)
	}
}

// TestResult_MarkFinished Test that if we mark a result as finished, that it says that it is finished
func TestResult_MarkFinished(t *testing.T) {
	// @TODO research how to prevent blocking in a chan during testing
}

// TestResult_MarkFailed Test that if we mark a Result as failed, then it returns a negative Success()
func TestResult_MarkFailed(t *testing.T) {
	res := NewResult()

	res.MarkFailed()
	if res.Success() {
		t.Error("Result that was marked failed, returned a positive success value")
	}
}

// TestResult_MarkSucceeded Test that if we mark a Result as successful, then it returns a positive Success()
func TestResult_MarkSucceeded(t *testing.T) {
	res := NewResult()

	res.MarkSucceeded()
	if !res.Success() {
		t.Error("Result that was marked succeeded, returned a negative success value")
	}
}

// TestResult_Merge Test that if we merge a Result into a Result, that it merges properly
func TestResult_Merge(t *testing.T) {
	res := NewResult()
	res.AddError(errors.New("one"))
	res.AddProperty(api.Property(&testPropertyOne{}))

	merge := NewResult()
	merge.AddError(errors.New("two"))
	merge.AddProperty(api.Property(&testPropertyTwo{}))

	res.Merge(api.Result(merge))

	errs := res.Errors()
	if len(errs) != 2 {
		t.Error("Result did not merge errors properly")
	}
	props := res.Properties()
	if len(props.Order()) != 2 {
		t.Error("Result did not merge properties properly")
	}
}
