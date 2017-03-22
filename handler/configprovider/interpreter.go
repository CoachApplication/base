package configprovider

import (
	"io"
	base_config "github.com/james-nesbitt/coach-base/operation/configuration"
)

type Interpreter interface {
	Get(source io.ReadCloser) (base_config.Config, error)
	Set(source base_config.Config) (io.ReadCloser, error)
}