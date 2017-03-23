package configprovider

import (
	base_config "github.com/james-nesbitt/coach-base/operation/configuration"
	"fmt"
)

type ScopedConfig struct {
	cMap map[string]base_config.Config
	cOrder []string
}

func NewScopedConfig() *ScopedConfig {
	return &ScopedConfig{}
}

// ScopedConfig Explicitly convert this struct to a config ScopedConfig interface
func (sc *ScopedConfig) ScopedConfig() base_config.ScopedConfig {
	return base_config.ScopedConfig(sc)
}

// Get a Config for a scope
func (sc *ScopedConfig) Get(scope string) (base_config.Config, error) {
	sc.safe()
	config, exists := sc.cMap[scope]
	if exists {
		return config, nil
	} else {
		return config, error(ConfigScopeNotFoundError{Scope: scope})
	}
}
// Set uses a passed Config to set a value to a scope
func (sc *ScopedConfig) Set(scope string, config base_config.Config) error {
	sc.safe()
	if _, exists := sc.cMap[scope]; !exists {
		sc.cOrder = append(sc.cOrder, scope)
	}
	sc.cMap[scope] = config
	return nil
}
// List available scopes
func (sc *ScopedConfig) List() []string {
	sc.safe()
	return sc.cOrder
}

func (sc *ScopedConfig) safe() {
	if &sc.cMap == nil {
		sc.cMap = map[string]base_config.Config{}
		sc.cOrder = []string{}
	}
}

type ConfigScopeNotFoundError struct {
	Scope string
}

func (csnfe ConfigNotHandlerdError) Error() string {
	return fmt.Sprintf("Config was not found at the reqyested scope %s", csnfe.Scope)
}
