package internal

var datastore *concurrentMapOfMaps = &concurrentMapOfMaps{}

func getDataStore() *concurrentMapOfMaps {
	return datastore
}
