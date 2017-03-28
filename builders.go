package standard

import (
	api "github.com/CoachApplication/coach-api"
)

type Builders struct {
	bMap   map[string]api.Builder
	bOrder []string
}

// NewBuilders Constructor for a new Builders
func NewBuilders() *Builders {
	return &Builders{}
}

// Builders Explicitly convert this to an api.Builders
func (sb *Builders) Builders() api.Builders {
	return api.Builders(sb)
}

// Add a new builder to the ordered set
func (sb *Builders) Add(b api.Builder) {
	sb.safe()
	key := b.Id()
	if _, found := sb.bMap[key]; !found {
		sb.bOrder = append(sb.bOrder, key)
	}
	sb.bMap[key] = b
}

// Get a builder that matches a key from the set
func (sb *Builders) Get(key string) (api.Builder, error) {
	sb.safe()
	if b, found := sb.bMap[key]; found {
		return b, nil
	} else {
		return b, error(&BuilderNotFoundError{key: key})
	}
}

// Order the builder keys from the set
func (sb *Builders) Order() []string {
	sb.safe()
	return sb.bOrder
}

// safe intitializer
func (sb *Builders) safe() {
	if sb.bOrder == nil {
		sb.bMap = map[string]api.Builder{}
		sb.bOrder = []string{}
	}
}

/**
 * Errors
 */

// BuilderNotFoundError for when a builder key does not exist in the list
type BuilderNotFoundError struct {
	key string
}

// Error string return (interface: error)
func (bnf *BuilderNotFoundError) Error() string {
	return "Builder not found: " + bnf.key
}
