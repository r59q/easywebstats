package persistence

import (
	"encoding/gob"
	"os"
	"r59q.com/easywebstats/internal/datastore"
)

var numberStoreFile = "data/nums.gob"

// Save to file
func SaveSerializedNumberStore(data *datastore.SerializedNumberStore) error {
	file, err := os.Create(numberStoreFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(data)
}

// Load from file
func ReadSerializedNumberStore() (*datastore.SerializedNumberStore, error) {
	file, err := os.Open(numberStoreFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data datastore.SerializedNumberStore
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&data)
	return &data, err
}
