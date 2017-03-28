package standard

import (
	"errors"
	"testing"

	api "github.com/CoachApplication/coach-api"
	"time"
)

// TestNewResult Test that the NewResult() function returns a struct that implements the api.Result interface
func TestNewResult(t *testing.T) {
	res := NewResult()

	var i interface{} = res
	if _, isResult := i.(api.Result); !isResult {
		t.Error("NewResult did not provide a struct which is a valid api.Result implementation", res)
	}
}

// TestResult_AddProperty test adding an retrieving properties to a Result
func TestResult_AddProperty(t *testing.T) {
	res := NewResult()

	newProp := NewTestProperty("test.1", "", "", "", nil).Property()
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
	res.AddProperty(NewTestProperty("test.1", "", "", "", nil).Property())

	merge := NewResult()
	merge.AddError(errors.New("two"))
	merge.AddProperty(NewTestProperty("test.2", "", "", "", nil).Property())

	res.Merge(merge.Result())

	errs := res.Errors()
	if len(errs) != 2 {
		t.Error("Result did not merge errors properly")
	}
	props := res.Properties()
	if len(props.Order()) != 2 {
		t.Error("Result did not merge properties properly", props.Order())
	}
}

// TestResult_Merge Test that if we merge a Result into a Result, that it the chans work properly
func TestResult_Merge_Finished(t *testing.T) { // use this delay to prevent timeout failures if our chans fail to close
	delay, _ := time.ParseDuration("1s")
	timeout := time.After(delay)
	now := time.Now()

	res := NewResult()

	merge1 := NewResult()
	merge2 := NewResult()

	res.Merge(merge1.Result())
	res.Merge(merge2.Result())

	fin := res.Finished()

	merge1.MarkFinished()
	merge2.MarkFinished()

	select {
	case <-fin:
		t.Log("Channel closed properly")
	case <-timeout:
		t.Error("Our opened finished channel failed to close after we finished all of the merged results.", time.Since(now))
	}

}

func TestResult_MarkFinished_Single(t *testing.T) {
	// use this delay to prevent timeout failures if our chans fail to close
	delay, _ := time.ParseDuration("1s")
	timeout := time.After(delay)
	now := time.Now()

	res := NewResult()
	fin := res.Finished()

	res.MarkFinished()

	select {
	case <-fin:
		t.Log("Channel closed properly")
	case <-timeout:
		t.Error("Our opened finished channel failed to close after we marked the result finished.", time.Since(now))
	}
}

func TestResult_MarkFinished_Multi(t *testing.T) { // use this delay to prevent timeout failures if our chans fail to close
	delay, _ := time.ParseDuration("1s")
	timeout := time.After(delay)
	now := time.Now()

	res := NewResult()

	pre1 := res.Finished()
	res.Finished() // ignore this one, we should still be safe
	pre2 := res.Finished()
	pre3 := res.Finished()

	t.Log("Marking finished :", res.MarkFinished())

	post1 := res.Finished()
	post2 := res.Finished()
	res.Finished() // ignore this one, we should still be safe
	post3 := res.Finished()

	t.Log("Marking finished :", res.MarkFinished())

	for i := 10; i > 0; i-- {
		select {

		case <-pre1:
			t.Log("Early Channel 1 closed properly")
		case <-pre3:
			t.Log("Early Channel 3 closed properly")
		case <-pre2:
			t.Log("Early Channel 2 closed properly")

		case <-post1:
			t.Log("Late Channel 1 closed properly")
		case <-post3:
			t.Log("Late Channel 3 closed properly")
		case <-post2:
			t.Log("Late Channel 2 closed properly")

		case <-timeout:
			t.Error("TIMEOUT", time.Since(now))
		}
	}
}
