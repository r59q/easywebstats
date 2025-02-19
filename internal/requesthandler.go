package internal

import "r59q.com/easywebstats/internal/datastore"

func HandleSetNum(name string, label string, value float64) float64 {
	store := datastore.GetNumberStore()
	store.Set(name, label, value)
	val := store.Get(name, label)
	return val
}

func HandleIncreaseNum(name string, label string, value float64) float64 {
	store := datastore.GetNumberStore()
	cur := store.Get(name, label)
	newVal := cur + value
	store.Set(name, label, newVal)
	return newVal
}

func HandleDecreaseNum(name string, label string, value float64) float64 {
	store := datastore.GetNumberStore()
	cur := store.Get(name, label)
	newVal := cur - value
	store.Set(name, label, newVal)
	return newVal
}

func ReadNumLabel(name string, label string) float64 {
	store := datastore.GetNumberStore()
	return store.Get(name, label)
}

func ReadNumName(name string) map[string]float64 {
	store := datastore.GetNumberStore()
	return store.GetLabels(name)
}
