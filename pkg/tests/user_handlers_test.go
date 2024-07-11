package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ceciivanov/go-challenge/pkg/handlers"
	"github.com/ceciivanov/go-challenge/pkg/models"
	"github.com/ceciivanov/go-challenge/pkg/repository"
	"github.com/gorilla/mux"
)

func TestGetUserFavorites(t *testing.T) {

	// Create a new Users Repository
	repo := repository.NewUsersRepository()
	repo.GenerateSampleUsers(5, 5)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/users/1/favorites", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := handlers.NewUserHandler(repo)

	// Create a new ResponseRecorder and assign the handler
	rr := httptest.NewRecorder()

	// Create a new router and assign the handler
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}/favorites", handler.GetUserFavorites).Methods("GET")
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetUserFavoritesUserNotFound(t *testing.T) {

	// Create a new Users Repository
	repo := repository.NewUsersRepository()
	repo.GenerateSampleUsers(5, 5)

	// Create a new HTTP request with an invalid user ID
	req, err := http.NewRequest("GET", "/users/999999/favorites", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := handlers.NewUserHandler(repo)

	// Create a new ResponseRecorder and assign the handler
	rr := httptest.NewRecorder()

	// Create a new router and assign the handler
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}/favorites", handler.GetUserFavorites).Methods("GET")
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	// Expected error message
	expectedErrorMsg := "User not found"
	if body := rr.Body.String(); !strings.Contains(body, expectedErrorMsg) {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedErrorMsg)
	}
}

func TestAddUserFavorite(t *testing.T) {

	// Create a new Users Repository
	repo := repository.NewUsersRepository()
	repo.GenerateSampleUsers(5, 5)

	// Prepare JSON payload for the asset you want to add as favorite
	payload := []byte(`{
		"id": 50,
		"type": "Insight",
		"description": "Sample Insight for testing",
		"text": "Testing Insight"
	}`)

	// Create a new HTTP request with the payload
	req, err := http.NewRequest("POST", "/users/1/favorites", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request Content-Type header
	req.Header.Set("Content-Type", "application/json")

	handler := handlers.NewUserHandler(repo)

	// Create a new ResponseRecorder and assign the handler
	rr := httptest.NewRecorder()

	// Create a new router and assign the handler
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}/favorites", handler.AddUserFavorite).Methods("POST")
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestAddUserFavoriteAssetExists(t *testing.T) {

	// Create a new Users Repository
	repo := repository.NewUsersRepository()

	// Initialize the Users Repository with one user and one favorite
	repo.Users = map[int]models.User{
		1: {
			ID: 1,
			Favourites: map[int]models.Asset{
				1: &models.Insight{
					ID:          1,
					Type:        models.InsightType,
					Description: "Sample Insight",
					Text:        "Sample Insight Text",
				},
			},
		},
	}

	// Define asset ID that already exists in the user's favorites
	payload := []byte(`{
		"id": 1,
		"type": "Insight",
		"description": "Sample Insight for testing",
		"text": "Testing Insight"
	}`)

	// Create a new HTTP request with the payload
	req, err := http.NewRequest("POST", "/users/1/favorites", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request Content-Type header
	req.Header.Set("Content-Type", "application/json")

	handler := handlers.NewUserHandler(repo)

	// Create a new ResponseRecorder and assign the handler
	rr := httptest.NewRecorder()

	// Create a new router and assign the handler
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}/favorites", handler.AddUserFavorite)
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestDeleteUserFavorite(t *testing.T) {

	// Create a new Users Repository
	repo := repository.NewUsersRepository()
	repo.GenerateSampleUsers(5, 5)

	// Create a new HTTP request
	req, err := http.NewRequest("DELETE", "/users/1/favorites/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := handlers.NewUserHandler(repo)

	// Create a new ResponseRecorder and assign the handler
	rr := httptest.NewRecorder()

	// Create a new router and assign the handler
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}/favorites/{assetID}", handler.DeleteUserFavorite).Methods("DELETE")
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify that the asset was removed from the user's favorites list
	// Example: Retrieve the updated user from the data store
	_, ok := repo.Users[1].Favourites[1]
	if ok {
		t.Errorf("expected asset ID 1 to be removed from user's favorites list, but it was found")
	}
}

func TestDeleteUserFavoriteAssetNotFound(t *testing.T) {

	// Create a new Users Repository
	repo := repository.NewUsersRepository()
	repo.GenerateSampleUsers(5, 5)

	// Create a new HTTP request
	req, err := http.NewRequest("DELETE", "/users/1/favorites/999999", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := handlers.NewUserHandler(repo)

	// Create a new ResponseRecorder and assign the handler
	rr := httptest.NewRecorder()

	// Create a new router and assign the handler
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}/favorites/{assetID}", handler.DeleteUserFavorite).Methods("DELETE")
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestEditUserFavorite(t *testing.T) {

	// Create a new Users Repository
	repo := repository.NewUsersRepository()
	repo.GenerateSampleUsers(5, 5)

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
		t.Fatal(err)
	}

	// Set the request Content-Type header
	req.Header.Set("Content-Type", "application/json")

	handler := handlers.NewUserHandler(repo)

	// Create a new ResponseRecorder and assign the handler
	rr := httptest.NewRecorder()

	// Create a new router and assign the handler
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}/favorites/{assetID}", handler.EditUserFavorite).Methods("PUT")
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestEditUserFavoriteNoIDMatch(t *testing.T) {

	// Create a new Users Repository
	repo := repository.NewUsersRepository()

	// Initialize the Users Repository with one user and one favorite
	repo.Users = map[int]models.User{
		1: {
			ID: 1,
			Favourites: map[int]models.Asset{
				1: &models.Insight{
					ID:          1,
					Type:        models.InsightType,
					Description: "Sample Insight",
					Text:        "Sample Insight Text",
				},
			},
		},
	}

	// Prepare JSON payload with an asset ID that does not match the URL
	payload := []byte(`{
		"id": 9999999,
		"type": "Insight",
		"description": "Updated Insight",
		"text": "Updated Insight Text"
	}`)

	// Create a new HTTP request with the payload
	req, err := http.NewRequest("PUT", "/users/1/favorites/1", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request Content-Type header
	req.Header.Set("Content-Type", "application/json")

	handler := handlers.NewUserHandler(repo)

	// Create a new ResponseRecorder and assign the handler
	rr := httptest.NewRecorder()

	// Create a new router and assign the handler
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}/favorites/{assetID}", handler.EditUserFavorite).Methods("PUT")
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestEditUserFavoriteNoTypeMatch(t *testing.T) {

	// Create a new Users Repository
	repo := repository.NewUsersRepository()

	// Initialize the Users Repository with one user and one favorite
	repo.Users = map[int]models.User{
		1: {
			ID: 1,
			Favourites: map[int]models.Asset{
				1: &models.Insight{
					ID:          1,
					Type:        models.InsightType,
					Description: "Sample Insight",
					Text:        "Sample Insight Text",
				},
			},
		},
	}

	// Prepare JSON payload with an asset type that does not match asset type in the data store
	payload := []byte(`{
		"id": 1,
		"type": "Chart",
		"description": "Updated Insight",
		"title": "Chart Updated",
		"xAxesTitle": "X-Updated",
		"yAxesTitle": "Y-Updated",
		"dataPoints": [
			{
			"X": 10,
			"Y": 10
			},
			{
			"X": 20,
			"Y": 20
			}
		]
	}`)

	// Create a new HTTP request with the payload
	req, err := http.NewRequest("PUT", "/users/1/favorites/1", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request Content-Type header
	req.Header.Set("Content-Type", "application/json")

	handler := handlers.NewUserHandler(repo)

	// Create a new ResponseRecorder and assign the handler
	rr := httptest.NewRecorder()

	// Create a new router and assign the handler
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}/favorites/{assetID}", handler.EditUserFavorite)
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
