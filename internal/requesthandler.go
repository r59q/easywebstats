package internal

func HandleSetNum(name string, label string, value float64) float64 {
	getDataStore().getOrCreateInnerMap(name).Set(label, value)
	val, _ := getDataStore().get(name, label)
	return val
}

func HandleIncreaseNum(name string, label string, value float64) float64 {
	store := getDataStore()
	val, exists := store.get(name, label)
	if !exists {
		store.set(name, label, value)
		return value
	}
	newValue := val + value
	store.set(name, label, newValue)
	return newValue
}

func HandleDecreaseNum(name string, label string, value float64) float64 {
	store := getDataStore()
	val, exists := store.get(name, label)
	if !exists {
		store.set(name, label, -value)
		return -value
	}
	newValue := val - value
	store.set(name, label, newValue)
	return newValue
}

func ReadNumLabel(name string, label string) (float64, bool) {
	store := getDataStore()
	return store.get(name, label)
}

func ReadNumName(name string) map[string]float64 {
	store := getDataStore()
	values := store.getOrCreateInnerMap(name).Values()

	return values
}
