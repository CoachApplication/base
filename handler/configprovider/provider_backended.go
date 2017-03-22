package configprovider

import (
	base_config "github.com/james-nesbitt/coach-base/operation/configuration"
)

type BackendConfigProvider struct {
	backends []Backend
}

func (bcp *BackendConfigProvider) Scopes() []string {
	scopes := []string{}
	for _, backend := range bcp.backends {
		scopes = append(scopes, backend.Scopes()...)
	}
	return scopes
}

func (bcp *BackendConfigProvider) Keys() []string {
	keys := []string{}
	for _, backend := range bcp.backends {
		keys = append(keys, backend.Keys()...)
	}
	return keys
}

func (bcp *BackendConfigProvider) Get(key, scope string) (base_config.Config, error) {
	for _, backend := range bcp.backends {
		if backend.Handles(key, scope) {
			return backend.Get(key, scope)
		}
	}
	return nil, error(ConfigNotHandlerdError{Key: key, Scope: scope})
}

func (bcp *BackendConfigProvider) Set(key, scope string, config base_config.Config) error {
	for _, backend := range bcp.backends {
		if backend.Handles(key, scope) {
			return backend.Set(key, scope, config)
		}
	}
	return error(ConfigNotHandlerdError{Key: key, Scope: scope})
}

type ConfigNotHandlerdError struct {
	Key   string
	Scope string
}

func (cnhe ConfigNotHandlerdError) Error() string {

}
