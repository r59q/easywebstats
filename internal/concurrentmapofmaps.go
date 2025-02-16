package internal

import (
	"sync"
)

type innerNumMap struct {
	mu   sync.RWMutex
	data map[string]float64
}

func newSafeInnerMap() *innerNumMap {
	return &innerNumMap{data: make(map[string]float64)}
}

func (m *innerNumMap) Set(key string, value float64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *innerNumMap) Get(key string) (float64, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, exists := m.data[key]
	return val, exists
}

func (m *innerNumMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func (m *innerNumMap) Values() map[string]float64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	// Clone and return map
	cloned := make(map[string]float64)
	for key, value := range m.data {
		cloned[key] = value
	}
	return cloned
}

type concurrentMapOfMaps struct {
	outerMap sync.Map
}

func (c *concurrentMapOfMaps) getOrCreateInnerMap(key string) *innerNumMap {
	actual, _ := c.outerMap.LoadOrStore(key, newSafeInnerMap())
	return actual.(*innerNumMap)
}

func (c *concurrentMapOfMaps) set(outerKey, innerKey string, value float64) {
	innerMap := c.getOrCreateInnerMap(outerKey)
	innerMap.Set(innerKey, value)
}

func (c *concurrentMapOfMaps) get(outerKey, innerKey string) (float64, bool) {
	if innerMap, ok := c.outerMap.Load(outerKey); ok {
		return innerMap.(*innerNumMap).Get(innerKey)
	}
	return 0, false
}

func (c *concurrentMapOfMaps) delete(outerKey, innerKey string) {
	if innerMap, ok := c.outerMap.Load(outerKey); ok {
		innerMap.(*innerNumMap).Delete(innerKey)
	}
}
