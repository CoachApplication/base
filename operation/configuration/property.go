package configuration

import (
	base_errors "github.com/james-nesbitt/coach-base/errors"
)

const(
	PROPERTY_KEY_KEY = "config.key"
	PROPERTY_KEY_KEYS = "config.keys"
	PROPERTY_KEY_SCOPE = "config.scope"
	PROPERTY_KEY_SCOPES = "config.scopes"
	PROPERTY_KEY_SCOPEDCONFIG = "config.scopedconfig"
)

// KeyProperty holds a string config key
type KeyProperty struct {

}

// Id Identify the property
func (kp *KeyProperty) Id() string {
	return PROPERTY_KEY_KEY
}

// KeysProperty holds a set of string config keys
type KeysProperty struct {

}

// Id Identify the property
func (kp *KeysProperty) Id() string {
	return PROPERTY_KEY_KEYS
}

// ScopeProperty holds a string config scope key
type ScopeProperty struct {

}

// Id Identify the property
func (sp *ScopeProperty) Id() string {
	return PROPERTY_KEY_SCOPE
}

// ScopesProperty holds a set of string config scope keys
type ScopesProperty struct {
	
}

// Id Identify the property
func (scp *ScopesProperty) Id() string {
	return PROPERTY_KEY_SCOPES
}

// ScopedConfigProperty is an api Property that holds a ScopedConfig struct
type ScopedConfigProperty struct {
	value ScopedConfig
}

// Id Identify the property
func (scp *ScopedConfigProperty) Id() string {
	return PROPERTY_KEY_SCOPEDCONFIG
}

// Id Identify the property
func (scp *ScopedConfigProperty) Type() string {
	return "coach.configuration.scopedconfig"
}

//
func (scp *ScopedConfigProperty) Get() interface{} {
	return scp.value
}

//
func (scp *ScopedConfigProperty) Set(value interface{}) error {
	if typedValue, success := value.(ScopedConfig); success {
		scp.value = typedValue
		return nil
	} else {
		return base_errors.PropertyWrongValueTypeError{Id: scp.Id(), ExpectedType: scp.Type()}
	}
}
