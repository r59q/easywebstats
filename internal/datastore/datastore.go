package datastore

type Datastore[V any] interface {
	Set(name string, label string, data V) V
	Get(name string, label string) V
	GetLabels(name string) map[string]V
}
