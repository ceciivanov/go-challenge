package service

import (
	"errors"
	"fmt"

	"github.com/ceciivanov/go-challenge/internal/models"
	"github.com/ceciivanov/go-challenge/internal/repository"
)

// UserService struct defines methods related to user operations
type UserService struct {
	UsersRepository *repository.UsersRepository
}

// NewUserService creates a new UserService instance
func NewUserService(repo *repository.UsersRepository) *UserService {
	return &UserService{
		UsersRepository: repo,
	}
}

// GetUserFavorites returns a map of user's favorite assets
func (s *UserService) GetUserFavorites(userID int) (map[int]models.Asset, error) {
	user, ok := s.UsersRepository.Users[userID]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user.Favourites, nil
}

// AddUserFavorite adds an asset to the user's favorites
func (s *UserService) AddUserFavorite(userID int, asset models.Asset) error {
	user, ok := s.UsersRepository.Users[userID]
	if !ok {
		return errors.New("user not found")
	}

	if _, ok := user.Favourites[asset.GetID()]; ok {
		return errors.New("asset already exists")
	}

	user.Favourites[asset.GetID()] = asset
	s.UsersRepository.Users[userID] = user
	return nil
}

// DeleteUserFavorite deletes an asset from the user's favorites
func (s *UserService) DeleteUserFavorite(userID, assetID int) error {
	user, ok := s.UsersRepository.Users[userID]
	if !ok {
		return errors.New("user not found")
	}

	if _, ok := user.Favourites[assetID]; !ok {
		return errors.New("asset not found")
	}

	delete(user.Favourites, assetID)
	s.UsersRepository.Users[userID] = user
	return nil
}

// EditUserFavorite edits an asset in the user's favorites
func (s *UserService) EditUserFavorite(userID int, assetID int, asset models.Asset) error {
	user, ok := s.UsersRepository.Users[userID]
	if !ok {
		return errors.New("user not found")
	}

	if _, ok := user.Favourites[assetID]; !ok {
		return errors.New("asset not found")
	}

	// Validate asset type and ID match the existing asset
	existingAsset := user.Favourites[assetID]

	if existingAsset.GetID() != asset.GetID() {
		return fmt.Errorf("edited asset ID does not match existing asset ID")
	}

	if existingAsset.GetType() != asset.GetType() {
		return fmt.Errorf("edited asset type does not match existing asset type")
	}

	user.Favourites[asset.GetID()] = asset
	s.UsersRepository.Users[userID] = user
	return nil
}
