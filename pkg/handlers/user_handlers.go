package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ceciivanov/go-challenge/pkg/data"
	"github.com/ceciivanov/go-challenge/pkg/models"
	"github.com/ceciivanov/go-challenge/pkg/utils"
	"github.com/gorilla/mux"
)

// Handler struct
type Handler struct {
	DataStore *data.DataStore
}

// NewHandler initializes and returns a new Handler
func NewHandler(ds *data.DataStore) *Handler {
	return &Handler{
		DataStore: ds,
	}
}

func (h *Handler) GetUserFavorites(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	// Check if the user exists
	user, ok := h.DataStore.Users[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return the user's favorites to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.Favourites)
}

func (h *Handler) AddUserFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	// Check if the user exists
	user, ok := h.DataStore.Users[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Read the request body into a byte slice
	newAssetData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Decode the asset data into the correct asset type
	newAsset, err := utils.DecodeAsset(newAssetData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user has a favourites map and create one if not
	if user.Favourites == nil {
		user.Favourites = make(map[int]models.Asset)
	}

	// Check if newAsset's id already exists in the map
	if _, ok := user.Favourites[newAsset.GetID()]; ok {
		http.Error(w, "Asset already exists", http.StatusBadRequest)
		return
	}

	// Add the new asset with the asset ID as the key to the user's favorites map and save the updated user
	user.Favourites[newAsset.GetID()] = newAsset
	h.DataStore.Users[userID] = user

	// Return the new asset to the client with a 201 Created status
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAsset)
}

func (h *Handler) DeleteUserFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	assetID, _ := strconv.Atoi(vars["assetID"])

	// Check if the user exists
	user, ok := h.DataStore.Users[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check if the asset exists in the user's favorites
	if _, ok := user.Favourites[assetID]; !ok {
		http.Error(w, "Asset not found in user's favorites", http.StatusNotFound)
		return
	}

	// Remove the asset from the user's favorites and save the updated user
	delete(user.Favourites, assetID)
	h.DataStore.Users[userID] = user

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) EditUserFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	assetID, _ := strconv.Atoi(vars["assetID"])

	// Check if the user exists
	user, ok := h.DataStore.Users[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check if the asset exists in the user's favortes
	asset, ok := user.Favourites[assetID]
	if !ok {
		http.Error(w, "Asset not found in user's favorites", http.StatusNotFound)
		return
	}

	// Read the request body into a byte slice
	updatedAssetData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Decode the asset data into the correct asset type
	updatedAsset, err := utils.DecodeAsset(updatedAssetData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate that assetID in the request body matches the assetID in the URL
	if updatedAsset.GetID() != assetID {
		http.Error(w, "Asset ID in the request body does not match the URL", http.StatusBadRequest)
		return
	}

	// Validate that the asset type in the request body matches the existing asset type
	if updatedAsset.GetType() != asset.GetType() {
		http.Error(w, "Asset type in the request body does not match the existing asset type", http.StatusBadRequest)
		return
	}

	// Update the asset in the user's favorites and save the updated user
	user.Favourites[assetID] = updatedAsset
	h.DataStore.Users[userID] = user

	// Return the updated asset to the client with a 200 OK status
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedAsset)
}
