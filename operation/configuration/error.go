package configuration

import (
	"fmt"
)

/**
 * Errors that the configuration system may create
 */

// ScopeNotFoundError is an error that indicates that a requested scope does not exist
type ScopeNotFoundError struct {
	// Scope string label for missing scope
	Scope string
}

// Error returns an error string (Error interface)
func (snf *ScopeNotFoundError) Error() string {
	return fmt.Sprintf("Scope not found : %s", snf.Scope)
}

// ScopeNoteWriteableError is an Error that says that the config scope cannot be written to
type ScopeNotWriteableError struct {}

// Error returns an error string (Error interface)
func (snw ScopeNotWriteableError) Error() string {
	return "Could not write to scope"
}

// ScopeEmpty error indicates that the retrieved scope has not data
type ScopeEmptyError struct {
	// Scope
	Scope string
}

// Error string outputter (interface: error)
func (see ScopeEmptyError) Error() string {
	return fmt.Sprintf("Scope had no data: %s", see.Scope)
}
