package standard

import (
	api "github.com/james-nesbitt/coach-api"
)

/**
 * Operation Usage implementations
 */

// ExternalOperationUsage Usage for an Operation that is meant to be used externally
type ExternalOperationUsage struct {
	Op api.Operation
}

// Allows return bool for if the Usage says that a certain key is allowed
func (eou *ExternalOperationUsage) Allows(key string) bool {
	val, _ := map[string]bool{
		api.UsageOperationPublicView: true,
		api.UsageOperationPublicExecute: true,
	}[key]
	return val
}

// InternalOperationUsage Usage for an Operation that is meant to be used internal
type InternalOperationUsage struct {
	Op api.Operation
}

// Allows return bool for if the Usage says that a certain key is allowed
func (iou *InternalOperationUsage) Allows(key string) bool {
	val, _ := map[string]bool{
		api.UsageOperationPublicView: false,
		api.UsageOperationPublicExecute: false,
	}[key]
	return val
}

/**
 * Property Usage implementations
 */

// InternalPropertyUsage Property Usage for a property that is only meant to be used internally
type InternalPropertyUsage struct {
	Prop api.Property
}

// Allows return bool for if the Usage says that a certain key is allowed
func (ipu *InternalPropertyUsage) Allows(key string) bool {
	val, _ := map[string]bool{
		api.UsagePropertyPublicView: false,
		api.UsagePropertyPublicWrite: false,
		api.UsagePropertyPublicRequired: false,
	}[key]
	return val
}

// ReadonlyPropertyUsage Property Usage for a Property that is only meant to be read
type ReadonlyPropertyUsage struct {
	Prop api.Property
}

// Allows return bool for if the Usage says that a certain key is allowed
func (rpu *ReadonlyPropertyUsage) Allows(key string) bool {
	val, _ := map[string]bool{
		api.UsagePropertyPublicView: true,
		api.UsagePropertyPublicWrite: false,
		api.UsagePropertyPublicRequired: false,
	}[key]
	return val
}

// OptionalPropertyUsage Property Usage for a Property that is allowed to be written to
type OptionalPropertyUsage struct {
	Prop api.Property
}

// Allows return bool for if the Usage says that a certain key is allowed
func (opu *OptionalPropertyUsage) Allows(key string) bool {
	val, _ := map[string]bool{
		api.UsagePropertyPublicView: true,
		api.UsagePropertyPublicWrite: true,
		api.UsagePropertyPublicRequired: false,
	}[key]
	return val
}

// RequiredPropertyUsage Property Usage for a Property which is expected to be written to before it's Operation is run
type RequiredPropertyUsage struct {
	Prop api.Property
}

// Allows return bool for if the Usage says that a certain key is allowed
func (rpu *RequiredPropertyUsage) Allows(key string) bool {
	val, _ := map[string]bool{
		api.UsagePropertyPublicView: true,
		api.UsagePropertyPublicWrite: true,
		api.UsagePropertyPublicRequired: true,
	}[key]
	return val
}
