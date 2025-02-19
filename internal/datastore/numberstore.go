package datastore

import "r59q.com/easywebstats/internal/concurrent"

type numberstore struct {
	numberMap *concurrent.Map[float64]
}

func (s numberstore) Set(name string, label string, data float64) float64 {
	s.numberMap.GetOrCreateInnerMap(name).Set(label, data)
	val := s.Get(name, label)
	return val
}

func (s numberstore) Get(name string, label string) float64 {
	value, wasSet := s.numberMap.GetOrCreateInnerMap(name).Get(label)
	if !wasSet {
		return 0
	}
	return value
}

func (s numberstore) GetLabels(name string) map[string]float64 {
	return s.numberMap.GetOrCreateInnerMap(name).Values()
}

var store = &numberstore{numberMap: &concurrent.Map[float64]{}}

func GetNumberStore() Datastore[float64] {
	return store
}
