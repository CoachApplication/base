package standard

import (
	api "github.com/james-nesbitt/coach-api"
	"fmt"
)

//Operations Operations implementation that maintains an ordered list
type Operations struct {
	oMap map[string]api.Operation
	oOrder []string
}

// NewOperations Constructor for Operations
func NewOperations() *Operations {
	return &Operations{}
}

// Operations Explicitly convert to an api.Operations interface
func (so *Operations) Operations() api.Operations {
	return api.Operations(so)
}

// safe lazy-initializer
func (so *Operations) safe() {
	if so.oMap == nil {
		so.oMap = map[string]api.Operation{}
		so.oOrder = []string{}
	}
}

/**
 * Interface: api.Operations
 */

// Add adds an operation to the list
func (so *Operations) Add(op api.Operation) error {
	so.safe()
	key := op.Id()
	if _, found := so.oMap[key]; !found {
		so.oOrder = append(so.oOrder, key)
	}
	so.oMap[key] = op
}

// Get retrieves an operation from the list by string key
func (so *Operations) Get(key string) (api.Operation, error) {
	so.safe()
	if op, found := so.oMap[key]; found {
		return op, nil
	} else {
		return op, error(OperationNotFound{Key: key})
	}
}

// List retrieves an ordered string list of operation keys
func (so *Operations) List() []string {
	so.safe()
	return so.oOrder
}

// Merge another set of Operation objects into this one
func (so *Operations) Merge(merge api.Operations) {
	for _, id := range merge.Order() {
		op, _ := merge.Get(id)
		so.Add(op)
	}
}

/**
 * Errors
 */

// OperationNotFound Error for when an operation cannot be found
type OperationNotFound struct {
	Key string
}

// Error Return error string
func (onf OperationNotFound) Error() string {
	return fmt.Sprintf("Operation was not found : %s", onf.Key)
}
