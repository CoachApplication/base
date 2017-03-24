package file

import (
	"io"
	"os"

	base_configprovider "github.com/james-nesbitt/coach-base/handler/configprovider"
)

// FileConnector is a ConfigProvider Connector that pulls config data from files.
// It requires only a struct which does the effort of mapping keys and scopes to file paths.
type FileConnector struct {
	filepaths FilePaths
}

func NewFileConnector(p FilePaths) *FileConnector {
	return &FileConnector{
		filepaths: p,
	}
}

func (fc *FileConnector) Connector() base_configprovider.Connector {
	return base_configprovider.Connector(fc)
}

func (fc *FileConnector) Scopes() []string {
	return fc.filepaths.Scopes()
}

func (fc *FileConnector) Keys() []string {
	return fc.filepaths.Keys()
}

func (fc *FileConnector) Get(key, scope string) (io.ReadCloser, error) {
	if path, err := fc.filepaths.Path(key, scope); err != nil {
		return nil, err
	} else {
		f, err := os.Open(path)
		return io.ReadCloser(f), err
	}
}

func (fc *FileConnector) Set(key, scope string, source io.ReadCloser) error {
	if path, err := fc.filepaths.Path(key, scope); err != nil {
		return err
	} else if file, err := os.Create(path); err != nil {
		return err
	} else {
		defer file.Close()
		defer source.Close()
		_, err := io.Copy(file, source)
		return err
	}
}
