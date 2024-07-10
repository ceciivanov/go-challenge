package main

import (
	"fmt"
	"net/http"

	"github.com/ceciivanov/go-challenge/pkg/data"
	"github.com/ceciivanov/go-challenge/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	data.GenerateMockData()

	// Create a new router from the Gorilla Mux package
	r := mux.NewRouter()

	// Define the routes

	// GET /users/{id}/favorites - Get a user's favorite assets
	r.HandleFunc("/users/{id}/favorites", handlers.GetUserFavorites).Methods("GET")

	// POST /users/{id}/favorites - Add a new asset to a user's favorites
	r.HandleFunc("/users/{id}/favorites", handlers.AddUserFavorite).Methods("POST")

	// DELETE /users/{id}/favorites/{assetID} - Remove an asset from a user's favorites
	r.HandleFunc("/users/{id}/favorites/{assetID}", handlers.DeleteUserFavorite).Methods("DELETE")

	// PUT /users/{id}/favorites/{assetID} - Edit an asset from a user's favorites
	r.HandleFunc("/users/{id}/favorites/{assetID}", handlers.EditUserFavorite).Methods("PUT")

	fmt.Println("Server is running on port 8080...")

	// Start the server
	http.ListenAndServe(":8080", r)
}
