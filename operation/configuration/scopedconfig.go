package configuration

type ScopedConfig struct {
	cMap map[string]Config 
	cOrder []string
}

// safe lazy-initializer
func (sc *ScopedConfig) safe() {
	if sc.cMap == nil {
		sc.cMap = map[string]Config{}
		sc.cOrder = []string{}
	}
}

func (sc *ScopedConfig) Get(scope string) (*Config, error) {
	sc.safe()
	if c, found := sc.cMap[scope]; found {
		return c, nil
	} else {
		return c, Error(&ConfigScopeNotFoundError{scope: scope})
	}
}

func (sc *ScopedConfig) Set(scope string, config *Config) error {
	sc.safe()
	if _, found := sc.cMap[scope]; !found {
		sc.cOrder = append(sc.cOrder, scope)
	}
	sc.cMap[scope] = config
}

func (sc *ScopedConfig) List() []string {
	sc.safe()
	return sc.cOrder
}

func (sc *ScopedConfig) load() (ScopedConfig, error) {
	
}

func (sc *ScopedConfig) save() error {
	
}
