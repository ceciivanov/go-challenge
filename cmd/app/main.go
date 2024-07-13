package main

import (
	"fmt"
	"net/http"

	"github.com/ceciivanov/go-challenge/internal/handlers"
	"github.com/ceciivanov/go-challenge/internal/repository"
	"github.com/ceciivanov/go-challenge/internal/service"
	"github.com/gorilla/mux"
)

func main() {

	NumberOfUsers := 10
	NumberOfAssets := 50

	// Create and initialize a new UsersRepository instance
	repo := repository.NewUsersRepository()
	repo.GenerateSampleUsers(NumberOfUsers, NumberOfAssets)

	// Create a new UserService instance
	userService := service.NewUserService(repo)

	// Create a new UserHandler instance
	userHandler := handlers.NewUserHandler(userService)

	// Create a new router from the Gorilla Mux package
	r := mux.NewRouter()

	// Register the routes
	userHandler.RegisterRoutes(r)

	// Start the server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}
