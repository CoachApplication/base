package standard

import (
	api "github.com/CoachApplication/coach-api"
	"fmt"
)

// Properties is a simple Properties implementation that keeps an ordered list of Properties
type Properties struct {
	pMap map[string]api.Property
	pOrder []string
}

// NewProperties is a Properties constructor
func NewProperties() *Properties {
	return &Properties{}
}

// Properties convert this object explicitly into a Properties interface
func (sp *Properties) Properties() api.Properties {
	return api.Properties(sp)
}

/**
 * Interface: api.Property
 */

// Add adds a new property to the list, keyed by property id
func (sp *Properties) Add(prop api.Property) {
	sp.safe()
	key := prop.Id()
	if _, found := sp.pMap[key]; !found {
		sp.pOrder = append(sp.pOrder, key)
	}
	sp.pMap[key] = prop
}

// Get retrieves a keyed property from the list
func (sp *Properties) Get(key string) (api.Property, error) {
	sp.safe()
	if prop, found := sp.pMap[key]; found {
		return prop, nil
	} else {
		return prop, error(PropertyNotFoundError{Key: key})
	}
}

// Order returns the ordered Property key list
func (sp *Properties) Order() []string {
	sp.safe()
	return sp.pOrder
}

// Merge Properties into this one
func (sp *Properties) Merge(merge api.Properties) {
	for _, id := range merge.Order() {
		prop, _ := merge.Get(id)
		sp.Add(prop)
	}
}

// Safe lazy initializer
func (sp *Properties) safe() {
	if sp.pMap == nil {
		sp.pMap = map[string]api.Property{}
		sp.pOrder = []string{}
	}
}

/**
 * Errors
 */

// PropertyNotFoundError error for when a Property could not be found by key
type PropertyNotFoundError struct {
	Key string
}

// Error return an error string (interface: error)
func (pnf PropertyNotFoundError) Error() string {
	return fmt.Sprintf("Property could not be found : %s", pnf.Key)
}
