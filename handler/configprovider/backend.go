package configprovider

import (
	base_config "github.com/james-nesbitt/coach-base/operation/configuration"
)

type Backend interface {
	Handles(key, scope string) bool
	Scopes() []string
	Keys() []string
	Get(key, scope string) (base_config.Config, error)
}
