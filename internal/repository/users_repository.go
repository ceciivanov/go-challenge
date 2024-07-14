package repository

import (
	"github.com/ceciivanov/go-challenge/internal/models"
	"github.com/ceciivanov/go-challenge/internal/repository/mock_data"
)

// UsersRepository is a struct that holds the users data in memory
type UsersRepository struct {
	Users map[int]models.User // map of models.User with user id as key
}

// NewUsersRepository creates a new UsersRepository instance
func NewUsersRepository() *UsersRepository {
	return &UsersRepository{
		Users: make(map[int]models.User),
	}
}

// GenerateSampleUsers generates sample users with sample assets
func (repo *UsersRepository) GenerateSampleUsers(NumberOfUsers, NumberOfAssets int) {
	repo.Users = mock_data.GenerateMockData(NumberOfUsers, NumberOfAssets)
}
