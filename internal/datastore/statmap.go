package datastore

import (
	"r59q.com/easywebstats/internal/concurrent"
	"time"
)

type ValueMetaData[V any] struct {
	CurrentValue     V
	HasPreviousValue bool
	PreviousValue    V
	LastUpdated      time.Time
}

type StatMapper[V any] interface {
	TemporalDatastore[V]
}

type StatMap[V any] struct {
	concurrentMap *concurrent.Map[ValueMetaData[V]]
}

func (sm *StatMap[V]) Get(name string, label string) (V, bool) {
	metaData, exists := sm.concurrentMap.GetOrCreateInnerMap(name).Get(label)
	if !exists {
		var zero V
		return zero, false
	}
	return metaData.CurrentValue, true
}

func (sm *StatMap[V]) Set(name string, label string, data V) V {
	previous, exists := sm.Get(name, label)
	if !exists {
		var zero V
		newMetaData := ValueMetaData[V]{
			CurrentValue:     data,
			HasPreviousValue: false,
			PreviousValue:    zero,
			LastUpdated:      time.Now(),
		}
		sm.concurrentMap.GetOrCreateInnerMap(name).Set(label, newMetaData)
		return data
	}
	newMetaData := ValueMetaData[V]{
		CurrentValue:     data,
		HasPreviousValue: true,
		PreviousValue:    previous,
		LastUpdated:      time.Now(),
	}
	sm.concurrentMap.GetOrCreateInnerMap(name).Set(label, newMetaData)
	return data
}

func (sm *StatMap[V]) GetPrevious(name string, label string) (V, bool) {
	val, exists := sm.concurrentMap.GetOrCreateInnerMap(name).Get(label)
	var zero V
	if !exists {
		return zero, false
	}
	if !val.HasPreviousValue {
		return zero, false
	}
	return val.PreviousValue, true
}

func (sm *StatMap[V]) GetLastUpdated(name string, label string) (time.Time, bool) {
	val, exists := sm.concurrentMap.GetOrCreateInnerMap(name).Get(label)
	if !exists {
		var zero time.Time
		return zero, false
	}

	return val.LastUpdated, true
}

func (sm *StatMap[V]) GetLabels(name string) map[string]V {
	keys := sm.concurrentMap.GetOrCreateInnerMap(name).Keys()
	values := sm.concurrentMap.GetOrCreateInnerMap(name).Values()

	resultMap := make(map[string]V)
	for _, key := range keys {
		resultMap[key] = values[key].CurrentValue
	}
	return resultMap
}

func CreateStatMap[V any]() StatMapper[V] {
	return &StatMap[V]{concurrentMap: &concurrent.Map[ValueMetaData[V]]{}}
}
