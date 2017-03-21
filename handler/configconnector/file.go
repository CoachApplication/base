package configconnector

import (
	"fmt"
	"io"
	"os"
	"path"
)

// FileScopePathsConfigConnector is a ConfigConnector that looks for a file in
// various scope labeled paths, where each scope corresponds to a particular
// path.
type FileScopePathsConfigConnector struct {
	filePathScopes map[string]FilePathScopes
}

// Config converts this YmlConfig to a Config interface (for clarity and validation)
func (fspcc *FileScopePathsConfigConnector) ConfigConnector() ConfigConnector {
	return ConfigConnector(fspcc)
}

// List lists all configs
func (fspcc *FileScopePathsConfigConnector) List() []string {
	keys := []string{}
	for key, _ := range fspcc.filePathScopes {
		keys = append(keys, key)
	}
	return keys
}

// ListScopes list all scopes for a config
func (fspcc *FileScopePathsConfigConnector) ListScopes(key string) []string {
	if scopePaths, found := fspcc.filePathScopes[key]; found {
		return scopePaths.ListScopes()
	} else {
		return []string{}
	}
}

// Get gets a configuration and apply it to a target struct
func (fspcc *FileScopePathsConfigConnector) Get(key string, scope string) (io.ReadCloser, error) {
	if filePathScopes, found := fspcc.filePathScopes[key]; !found {
		return nil, error(ConfigNotFoundError{Key: key})
	} else if scopeFile, err := filePathScopes.GetFilePath(scope); err != nil {
		return nil, err
	} else {
		file, err := os.Open(scopeFile)
		return io.ReadCloser(file), err
	}
}

// Set sets a Config value by converting a passed struct into a configuration
func (fspcc *FileScopePathsConfigConnector) Set(key string, scope string, source io.ReadCloser) error {
	if filePathScopes, found := fspcc.filePathScopes[key]; !found {
		return error(ConfigNotFoundError{})
	} else if scopeFile, err := filePathScopes.GetFilePath(scope); err != nil {
		return err
	} else if file, err := os.Create(scopeFile); err != nil {
		return err
	} else {
		defer file.Close()
		defer source.Close()
		_, err := io.Copy(file, source)
		return err
	}
}

/**
 * Paths ordered set
 */

// FilePathScopes An ordered list of scope keys, with corresponding file paths
type FilePathScopes struct {
	filename string
	pMap     map[string]string
	pOrder   []string
}

// Construct for FileScopePathsConfigConnector with a filename filename that
// starts with empty paths list
func NewFilePathScopes(filename string) *FilePathScopes {
	return &FilePathScopes{filename: filename}
}

// lazy initializer
func (fps *FilePathScopes) safe() {
	if fps.pMap == nil {
		fps.pMap = map[string]string{}
		fps.pOrder = []string{}
	}
}

// List all scope keys
func (fps *FilePathScopes) ListScopes() []string {
	fps.safe()
	return fps.pOrder
}

// Get a path for a scope key
func (fps *FilePathScopes) GetFilePath(scope string) (string, error) {
	fps.safe()
	if p, found := fps.pMap[scope]; found {
		return path.Join(p, fps.filename), nil
	} else {
		return "", error(ScopeNotFoundError{Scope: scope})
	}
}

// Set a path for a scope key
func (fps *FilePathScopes) SetPathScope(scope string, path string) {
	if _, found := fps.pMap[scope]; !found {
		fps.pOrder = append(fps.pOrder, scope)
	}
	fps.pMap[scope] = path
}

/**
 * Errors used
 */

// NoFileError indicated that no file was loaded
type NoFileError struct {
	Path string
}

// Error returns an error string (interface: error)
func (nfe NoFileError) Error() string {
	return fmt.Sprintf("No File was found at path : %s", nfe.Path)
}

// ConfigNotFoundError Config was not found Error
type ConfigNotFoundError struct {
	Key string
}

// Error returns an error string (interface: error)
func (cnfe ConfigNotFoundError) Error() string {
	return fmt.Sprintf("Config was not found: %s", cnfe.Key)
}
