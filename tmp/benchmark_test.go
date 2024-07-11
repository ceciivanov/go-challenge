package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ceciivanov/go-challenge/pkg/handlers"
	"github.com/ceciivanov/go-challenge/pkg/repository"
	"github.com/gorilla/mux"
)

const (
	numUsers     = 1000
	numFavorites = 100
)

func BenchmarkGetUserFavorites(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Create a new Users Repository
		repo := repository.NewUsersRepository()
		repo.GenerateSampleUsers(numUsers, numFavorites)

		// Create a new HTTP request
		req, err := http.NewRequest("GET", "/users/1/favorites", nil)
		if err != nil {
			b.Fatal(err)
		}

		handler := handlers.NewUserHandler(repo)

		// Create a new router and assign the handler
		r := mux.NewRouter()
		r.HandleFunc("/users/{id}/favorites", handler.GetUserFavorites).Methods("GET")

		rr := httptest.NewRecorder()
		start := time.Now()
		r.ServeHTTP(rr, req)
		elapsed := time.Since(start)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			b.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// Log the time taken for the request
		b.Logf("Request took %v", elapsed)
	}
}

func BenchmarkAddUserFavorite(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Create a new Users Repository
		repo := repository.NewUsersRepository()
		repo.GenerateSampleUsers(numUsers, numFavorites)

		// Prepare JSON payload for the asset you want to add as favorite
		payload := []byte(`{
			"id": 5001,
			"type": "Insight",
			"description": "Sample Insight for testing",
			"text": "Testing Insight"
		}`)

		// Create a new HTTP request with the payload
		req, err := http.NewRequest("POST", "/users/1/favorites", bytes.NewBuffer(payload))
		if err != nil {
			b.Fatal(err)
		}

		// Set the request Content-Type header
		req.Header.Set("Content-Type", "application/json")

		handler := handlers.NewUserHandler(repo)

		// Create a new router and assign the handler
		r := mux.NewRouter()
		r.HandleFunc("/users/{id}/favorites", handler.AddUserFavorite).Methods("POST")

		rr := httptest.NewRecorder()
		start := time.Now()
		r.ServeHTTP(rr, req)
		elapsed := time.Since(start)

		// Check the status code
		if status := rr.Code; status != http.StatusCreated {
			b.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}

		// Log the time taken for the request
		b.Logf("Request took %v", elapsed)
	}
}

func BenchmarkDeleteUserFavorite(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Create a new Users Repository
		repo := repository.NewUsersRepository()
		repo.GenerateSampleUsers(numUsers, numFavorites)

		// Create a new HTTP request
		req, err := http.NewRequest("DELETE", "/users/1000/favorites/100", nil)
		if err != nil {
			b.Fatal(err)
		}

		handler := handlers.NewUserHandler(repo)

		// Create a new router and assign the handler
		r := mux.NewRouter()
		r.HandleFunc("/users/{id}/favorites/{assetID}", handler.DeleteUserFavorite).Methods("DELETE")

		rr := httptest.NewRecorder()
		start := time.Now()
		r.ServeHTTP(rr, req)
		elapsed := time.Since(start)

		// Check the status code
		if status := rr.Code; status != http.StatusNoContent {
			b.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
		}

		// Log the time taken for the request
		b.Logf("Request took %v", elapsed)
	}
}

func BenchmarkEditUserFavorite(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Create a new Users Repository
		repo := repository.NewUsersRepository()
		repo.GenerateSampleUsers(numUsers, numFavorites)

		// Prepare JSON payload for the asset you want to update
		payload := []byte(`{
			"id": 1,
			"type": "Insight",
			"description": "Updated Insight",
			"text": "Updated Insight Text"
		}`)

		// Create a new HTTP request with the payload
		req, err := http.NewRequest("PUT", "/users/1/favorites/1", bytes.NewBuffer(payload))
		if err != nil {
			b.Fatal(err)
		}

		// Set the request Content-Type header
		req.Header.Set("Content-Type", "application/json")

		handler := handlers.NewUserHandler(repo)

		// Create a new router and assign the handler
		r := mux.NewRouter()
		r.HandleFunc("/users/{id}/favorites/{assetID}", handler.EditUserFavorite).Methods("PUT")

		rr := httptest.NewRecorder()
		start := time.Now()
		r.ServeHTTP(rr, req)
		elapsed := time.Since(start)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			b.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// Log the time taken for the request
		b.Logf("Request took %v", elapsed)
	}
}
