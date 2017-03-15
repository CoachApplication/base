package configconnector


/**
 * Errors
 */

// KeyNotFoundError Error for when a config was not found for a config key
type KeyNotFoundError struct {
	Key string	
}

// Error string output for error (interface: Error)
func (knfe KeyNotFoundError) Error() string {
	fmt.Sprintf("The config key '%s' was not found", knfe.Key)
}

// ScopeNotFoundError Error for when a config scope was not found for a config key
type ScopeNotFoundError struct {
	Key string
	Scope string
}

// Error string output for error (interface: Error)
func (snfe ScopeNotFoundError) Error() string {
	fmt.Sprintf("The config scope '%s' was not found for config `%s`", snfe.Key, snfe.Scope)
}

// SetError Error for when a connector failed to save
type SetError struct {
	Key string
	Scope string
}

// Error string output for error (interface: Error)
func (se SetError) Error() string {
	fmt.Sprintf("The config scope '%s' was not saved for config `%s`", ge.Key, ge.Scope)
}

// GetError Error for when a config fails to retrieve
type GetError struct {
	Key string
	Scope string
}

// Error string output for error (interface: Error)
func (ge GetError) Error() string {
	fmt.Sprintf("The config scope '%s' was not retrieved for config `%s`", ge.Key, ge.Scope)
}
