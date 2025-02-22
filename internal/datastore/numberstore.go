package datastore

type NumberStore interface {
	Set(name string, label string, data float64) float64
	Get(name string, label string) float64
	GetLabels(name string) map[string]float64
}

type numberStore struct {
	numberMap StatMapper[float64]
}

func (s *numberStore) Set(name string, label string, data float64) float64 {
	newValue := s.numberMap.Set(name, label, data)
	return newValue
}

func (s *numberStore) Get(name string, label string) float64 {
	val, exists := s.numberMap.Get(name, label)
	if !exists {
		return 0
	}
	return val
}

func (s *numberStore) GetLabels(name string) map[string]float64 {
	return s.numberMap.GetLabels(name)
}

var store = &numberStore{numberMap: CreateStatMap[float64]()}

func GetNumberStore() NumberStore {
	return store
}
