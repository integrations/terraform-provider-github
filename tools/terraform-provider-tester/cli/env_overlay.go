package cli

import "sync"

type envOverlay struct {
	mu        sync.RWMutex
	overrides map[string]string
	baseGet   func(string) string
	baseSet   func(string, string) error
}

func newEnvOverlay(baseGet func(string) string, baseSet func(string, string) error) *envOverlay {
	if baseGet == nil {
		baseGet = func(string) string { return "" }
	}
	if baseSet == nil {
		baseSet = func(string, string) error { return nil }
	}
	return &envOverlay{
		overrides: make(map[string]string),
		baseGet:   baseGet,
		baseSet:   baseSet,
	}
}

func (e *envOverlay) Get(key string) string {
	e.mu.RLock()
	val, ok := e.overrides[key]
	e.mu.RUnlock()
	if ok {
		return val
	}
	return e.baseGet(key)
}

func (e *envOverlay) Set(key, val string) error {
	if err := e.baseSet(key, val); err != nil {
		return err
	}
	e.mu.Lock()
	e.overrides[key] = val
	e.mu.Unlock()
	return nil
}
