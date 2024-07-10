package main

import (
	"fmt"
	"net/http"

	"github.com/ceciivanov/go-challenge/pkg/data"
	"github.com/ceciivanov/go-challenge/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	numberOfUsers := 1
	numberOfAssets := 2

	// Create a new DataStore and generate mock data
	ds := data.NewDataStore()
	ds.GenerateMockDataStore(numberOfUsers, numberOfAssets)

	// Create a new Handler instance with the DataStore
	handler := handlers.NewHandler(ds)

	// Create a new router from the Gorilla Mux package
	r := mux.NewRouter()

	// Define the routes using handler methods

	// GET /users/{id}/favorites - Get a user's favorite assets
	r.HandleFunc("/users/{id}/favorites", handler.GetUserFavorites).Methods("GET")

	// POST /users/{id}/favorites - Add a new asset to a user's favorites
	r.HandleFunc("/users/{id}/favorites", handler.AddUserFavorite).Methods("POST")

	// DELETE /users/{id}/favorites/{assetID} - Remove an asset from a user's favorites
	r.HandleFunc("/users/{id}/favorites/{assetID}", handler.DeleteUserFavorite).Methods("DELETE")

	// PUT /users/{id}/favorites/{assetID} - Edit an asset from a user's favorites
	r.HandleFunc("/users/{id}/favorites/{assetID}", handler.EditUserFavorite).Methods("PUT")

	fmt.Println("Server is running on port 8080...")

	// Start the server
	http.ListenAndServe(":8080", r)
}
