package configuration

import (
	"io"
)

/**
 *  ConfigConnector is a single backend that provides enough methods to provide
 * all three config operations
 */
type Connector interface {
	// List lists all configs
	List() []string
	// ListScopes list all scopes for a config
	ListScopes(key string) []string

	// Get retrieves ScopedConfig for the Config Get operation
	Get(key string, scope string) (io.ReadCloser, error)
	// Set Pushes provide ScopedConfig to a
	Set(key string, scope string, readcloser io.ReadCloser) error
}
