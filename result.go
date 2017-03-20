package standard

import (
	api "github.com/james-nesbitt/coach-api"
)

// MakeSuccessfulResult Quickly make an api.Result which is already finished and marked successful
func MakeSuccessfulResult() api.Result {
	res := NewResult()
	res.MarkSucceeded()
	res.MarkFinished()
	return res.Result()
}

// Result is as default Result implementation
type Result struct {
	errors  []error
	success bool

	finished []chan bool

	properties Properties

	waiting int // How many merged results have open finished channels
}

// NewResult constructs a new Result
func NewResult() *Result {
	return &Result{
		errors:     []error{},
		success:    true, // default to successful
		finished:   []chan bool{},
		properties: *NewProperties(),
		waiting:    0,
	}
}

// Result explicitly converts this struct to the Result interface (for clarity and validation)
func (sr *Result) Result() api.Result {
	return api.Result(sr)
}

/**
 * Result interface
 */

// Return a slice of any errors that occurred
func (sr *Result) Errors() []error {
	return sr.errors
}

// Finished returns a tracking bool channel that can be used to mark when the operation is completed
func (sr *Result) Finished() chan bool {
	finished := make(chan bool)
	sr.finished = append(sr.finished, finished)
	return finished
}

// Success returns a boolean success value
func (sr *Result) Success() bool {
	return sr.success
}

// Properties returns an ordered list of property values for the result
func (sr *Result) Properties() api.Properties {
	return sr.properties.Properties()
}

/**
 * Methods for creating the result data
 */

// AddError adds an Error to the result
func (sr *Result) AddError(err error) {
	sr.errors = append(sr.errors, err)
}

// AddProperty adds a Property to the result
func (sr *Result) AddProperty(prop api.Property) {
	sr.properties.Add(prop)
}

// MarkFailed marks this result as failed
func (sr *Result) MarkFailed() {
	sr.success = false
}

// MarkSucceeded marks this result as succeeded
func (sr *Result) MarkSucceeded() {
	sr.success = true
}

// MarkFinished marks this result operations as completed
func (sr *Result) MarkFinished() {
	finishedList := sr.finished
	sr.finished = []chan bool{}
	go func(finishedList []chan bool) {
		for _, eachFinished := range finishedList {
			eachFinished <- true
			close(eachFinished)
		}
	}(finishedList)
}

// Merge a result into this result
func (sr *Result) Merge(merge api.Result) {
	//sr.waiting = sr.waiting + 1
	//go func() {
	//	<-merge.Finished()
	//	sr.waiting = sr.waiting - 1
	//	sr.MarkFinished()
	//}()

	if !merge.Success() {
		sr.MarkFailed()
	}

	sr.properties.Merge(merge.Properties())

	for _, err := range merge.Errors() {
		sr.AddError(err)
	}
}
