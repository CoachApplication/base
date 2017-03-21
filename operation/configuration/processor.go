package configuration

import "io"


type Processor interface {
	Get(io.ReadCloser) (Config, error)
	Set(Config, io.ReadCloser) error
}
