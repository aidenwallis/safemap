package safemap

import "sync"

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

// Create a new Safemap instance
func New() *SafeMap {
	m := &SafeMap{
		data: map[string]interface{}{},
	}
	return m
}

// Set set a value into the hashmap.
func (m *SafeMap) Set(key string, value interface{}) {
	m.mu.Lock()
	m.data[key] = value
	m.mu.Unlock()
}

// Get retrieve a value from the hashmap
func (m *SafeMap) Get(key string) (interface{}, bool) {
	m.mu.RLock()
	val, ok := m.data[key]
	m.mu.RUnlock()
	return val, ok
}

// Has check if hashmap has a key
func (m *SafeMap) Has(key string) bool {
	m.mu.RLock()
	_, ok := m.data[key]
	m.mu.RUnlock()
	return ok
}

// Count get an integer of the number of keys in the hashmap
func (m *SafeMap) Count() int {
	m.mu.RLock()
	count := len(m.data)
	m.mu.RUnlock()
	return count
}

// Delete remove a key from the hashmap
func (m *SafeMap) Delete(key string) {
	m.mu.Lock()
	delete(m.data, key)
	m.mu.Unlock()
}
