package main

import (
	"r59q.com/easywebstats/api"
	"r59q.com/easywebstats/internal/datastore"
	"r59q.com/easywebstats/internal/persistence"
	"time"
)

func main() {
	fileData, err := persistence.ReadSerializedNumberStore()
	if err == nil {
		datastore.LoadSerializedNumberStore(fileData)
		println("Successfully loaded number store from file")
	} else {
		println("Couldn't load store from file:", err.Error())
	}

	go func() {
		ticker := time.NewTicker(600 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				store := datastore.GetNumberStore()
				err := persistence.SaveSerializedNumberStore(store.Serialize())
				if err != nil {
					println("Error saving data:", err.Error())
				}
				println("Flushed to filesystem")
			}
		}
	}()

	api.RunGinSever()
}
