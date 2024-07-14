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

	// Create UserService and Handler for it
	userService := service.NewUserService(repo)
	userHandler := handlers.NewUserHandler(userService)

	// Create a new router from the Gorilla Mux package and register the respective routes for the userHandler
	r := mux.NewRouter()
	userHandler.RegisterRoutes(r)

	// Start the server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}
