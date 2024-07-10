package data

import "github.com/ceciivanov/go-challenge/pkg/models"

// DataStore is a struct that holds the users data in memory
type DataStore struct {
	Users map[int]models.User
}

// NewDataStore creates a new DataStore instance
func NewDataStore() *DataStore {
	return &DataStore{
		Users: make(map[int]models.User),
	}
}

func (ds *DataStore) GenerateMockDataStore(numberOfUsers, numberOfAssets int) {
	// Mock data generation logic
	ds.Users = GenerateMockData(numberOfUsers, numberOfAssets)
}
