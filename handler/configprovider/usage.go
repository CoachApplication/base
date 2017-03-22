package configprovider

type BackendUsage interface {
	Handles(key, scope string) bool
}

type AllBackendUsage struct{}

func (abu *AllBackendUsage) Handles(key, scope string) bool {
	return true
}
