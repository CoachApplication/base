package configuration

// Config encapsulation of configuration
type Config interface {
	// Marshall gets a configuration and apply it to a target struct
	Get(interface{}) error
	// UnMarshall sets a Config value by converting a passed struct into a configuration
	// The expects that the values assigned are permanently saved
	Set(interface{}) error
}
