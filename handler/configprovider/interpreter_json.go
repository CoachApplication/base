package configprovider

import (
	"encoding/json"
	"io"

	base_config "github.com/james-nesbitt/coach-base/operation/configuration"
)

type JsonInterpreter struct{}

func (ji *JsonInterpreter) Get(source io.ReadCloser) (base_config.Config, error) {

}
func (ji *JsonInterpreter) Set(source interface{}) (io.ReadCloser, error) {

}

type JsonConfig struct {
	interpreter JsonInterpreter
}

// Marshall gets a configuration and apply it to a target struct
func (jc *JsonConfig) Get(target interface{}) error {
	//if rc, err := jc.interpreter.Get
	//d := json.NewDecoder()
	//return d.Decode(target)
}

// UnMarshall sets a Config value by converting a passed struct into a configuration
// The expects that the values assigned are permanently saved
func (jc *JsonConfig) Set(source interface{}) error {

}

type jsonReadCloser struct {
	reader io.Reader
}

func (jrc jsonReadCloser) Read(p []byte) (n int, err error) {
	reader, err := jrc.reader.Read(p)
	return reader, err
}

func (jrc jsonReadCloser) Close() error {
	return nil
}
