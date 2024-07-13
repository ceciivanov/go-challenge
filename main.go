package main

import (
	"fmt"
	"net/http"

	"github.com/ceciivanov/go-challenge/pkg/handlers"
	"github.com/ceciivanov/go-challenge/pkg/repository"
	"github.com/gorilla/mux"
)

func main() {

	// Create and initialize a new UsersRepository instance
	NumberOfUsers := 10
	NumberOfAssets := 50

	repo := repository.NewUsersRepository()
	repo.GenerateSampleUsers(NumberOfUsers, NumberOfAssets)

	// Create a new Handler instance and pass the repository to it
	handler := handlers.NewUserHandler(repo)

	// Create a new router from the Gorilla Mux package
	r := mux.NewRouter()

	// Define the routes using handler methods
	r.HandleFunc("/users/{id}/favorites", handler.GetUserFavorites).Methods("GET")
	r.HandleFunc("/users/{id}/favorites", handler.AddUserFavorite).Methods("POST")
	r.HandleFunc("/users/{id}/favorites/{assetID}", handler.DeleteUserFavorite).Methods("DELETE")
	r.HandleFunc("/users/{id}/favorites/{assetID}", handler.EditUserFavorite).Methods("PUT")

	// Start the server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}
