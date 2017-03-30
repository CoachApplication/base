package base

import (
	api "github.com/CoachApplication/api"
)

// UsageDecoratingProperty a Property decorator that overrides Usage with a provided Usage
type UsageDecoratingProperty struct {
	prop  api.Property
	usage api.Usage
}

func UsageDecorateProperty(prop api.Property, usage api.Usage) api.Property {
	return (&UsageDecoratingProperty{prop: prop, usage: usage}).Property()
}

func (udp *UsageDecoratingProperty) Property() api.Property {
	return api.Property(udp)
}

func (udp *UsageDecoratingProperty) Id() string {
	return udp.prop.Id()
}
func (udp *UsageDecoratingProperty) Type() string {
	return udp.prop.Type()
}
func (udp *UsageDecoratingProperty) Ui() api.Ui {
	return udp.prop.Ui()
}
func (udp *UsageDecoratingProperty) Usage() api.Usage {
	return udp.usage
}
func (udp *UsageDecoratingProperty) Validate() bool {
	return udp.prop.Validate()
}
func (udp *UsageDecoratingProperty) Get() interface{} {
	return udp.prop.Get()
}
func (udp *UsageDecoratingProperty) Set(val interface{}) error {
	return udp.prop.Set(val)
}
