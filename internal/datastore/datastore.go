package datastore

import "time"

type DataStore[V any] interface {
	Set(name string, label string, data V) V
	Get(name string, label string) (V, bool)
	GetLabels(name string) map[string]V
	GetNames() []string
}

type TemporalDatastore[V any] interface {
	DataStore[V]
	GetPrevious(name string, label string) (V, bool)
	GetLastUpdated(name string, label string) (time.Time, bool)
}
