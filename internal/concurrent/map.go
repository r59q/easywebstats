package concurrent

import (
	"sync"
)

type InnerMap[V any] struct {
	mu   sync.RWMutex
	data map[string]V
}

func newSafeInnerMap[V any]() *InnerMap[V] {
	return &InnerMap[V]{data: make(map[string]V)}
}

func (m *InnerMap[V]) Set(key string, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *InnerMap[V]) Get(key string) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, exists := m.data[key]
	return val, exists
}

func (m *InnerMap[V]) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func (m *InnerMap[V]) Values() map[string]V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	// Clone and return map
	cloned := make(map[string]V)
	for key, value := range m.data {
		cloned[key] = value
	}
	return cloned
}

type Map[V any] struct {
	outerMap sync.Map
}

func (c *Map[V]) GetOrCreateInnerMap(key string) *InnerMap[V] {
	actual, _ := c.outerMap.LoadOrStore(key, newSafeInnerMap[V]())
	return actual.(*InnerMap[V])
}

func (c *Map[V]) set(outerKey, innerKey string, value V) {
	inner := c.GetOrCreateInnerMap(outerKey)
	inner.Set(innerKey, value)
}

func (c *Map[V]) get(outerKey, innerKey string) (V, bool) {
	if actual, ok := c.outerMap.Load(outerKey); ok {
		return actual.(*InnerMap[V]).Get(innerKey)
	}
	var zero V
	return zero, false
}

func (c *Map[V]) delete(outerKey, innerKey string) {
	if actual, ok := c.outerMap.Load(outerKey); ok {
		actual.(*InnerMap[V]).Delete(innerKey)
	}
}
