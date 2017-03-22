package configprovider

import (
	base_config "github.com/james-nesbitt/coach-base/operation/configuration"
)

type ConnectorInterpreterBackend struct {
	usage       BackendUsage
	connector   Connector
	interpreter Interpreter
}

func NewConnectorInterpreterBackend(usage BackendUsage, connector Connector, interpreter Interpreter) *ConnectorInterpreterBackend {
	return &ConnectorInterpreterBackend{
		usage:       usage,
		connector:   connector,
		interpreter: interpreter,
	}
}

func (cib *ConnectorInterpreterBackend) Handles(key, scope string) bool {
	return cib.usage.Handles(key, scope)
}

func (cib *ConnectorInterpreterBackend) Scopes() []string {
	return cib.connector.Scopes()
}

func (cib *ConnectorInterpreterBackend) Keys() []string {
	return cib.connector.Scopes()
}

func (cib *ConnectorInterpreterBackend) Get(key, scope string) (base_config.Config, error) {
	if source, err := cib.connector.Get(key, scope); err != nil {
		return nil, err
	} else if conf, err := cib.interpreter.Get(source); err != nil {
		return nil, err
	} else {
		return conf, nil
	}
}

func (cib *ConnectorInterpreterBackend) Set(key, scope string, config base_config.Config) error {
	if source, err := cib.interpreter.Set(config); err != nil {
		return err
	} else if err := cib.connector.Set(key, scope, source); err != nil {
		return err
	} else {
		return nil
	}
}
