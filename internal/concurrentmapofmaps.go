package internal

import (
	"sync"
)

type innerMap[V any] struct {
	mu   sync.RWMutex
	data map[string]V
}

type concurrentMapOfMaps struct {
	outerMap sync.Map
}

func newSafeInnerMap() *innerMap[float64] {
	return &innerMap[float64]{data: make(map[string]float64)}
}

func (m *innerMap[float64]) Set(key string, value float64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *innerMap[float64]) Get(key string) (float64, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, exists := m.data[key]
	return val, exists
}

func (m *innerMap[float64]) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func (m *innerMap[float64]) Values() map[string]float64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	// Clone and return map
	cloned := make(map[string]float64)
	for key, value := range m.data {
		cloned[key] = value
	}
	return cloned
}

func (c *concurrentMapOfMaps) getOrCreateInnerMap(key string) *innerMap[float64] {
	actual, _ := c.outerMap.LoadOrStore(key, newSafeInnerMap())
	return actual.(*innerMap[float64])
}

func (c *concurrentMapOfMaps) set(outerKey, innerKey string, value float64) {
	inner := c.getOrCreateInnerMap(outerKey)
	inner.Set(innerKey, value)
}

func (c *concurrentMapOfMaps) get(outerKey, innerKey string) (float64, bool) {
	if actual, ok := c.outerMap.Load(outerKey); ok {
		return actual.(*innerMap[float64]).Get(innerKey)
	}
	return 0, false
}

func (c *concurrentMapOfMaps) delete(outerKey, innerKey string) {
	if actual, ok := c.outerMap.Load(outerKey); ok {
		actual.(*innerMap[float64]).Delete(innerKey)
	}
}
