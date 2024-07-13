package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ceciivanov/go-challenge/pkg/models"
	"github.com/ceciivanov/go-challenge/pkg/repository"
	"github.com/gorilla/mux"
)

type testCase struct {
	name           string
	method         string
	url            string
	payload        interface{}
	expectedStatus int
	expectedBody   string
}

func runTestCase(t *testing.T, tc testCase) {
	// Initialize the repository with 3 users and 3 favorites each
	repo := repository.NewUsersRepository()
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
				2: &models.Chart{
					ID:          2,
					Type:        models.ChartType,
					Description: "Sample Chart",
					Title:       "Sample Chart Title",
					XAxesTitle:  "X-Axis",
					YAxesTitle:  "Y-Axis",
					DataPoints: []models.Point{
						{X: 10, Y: 10},
						{X: 20, Y: 20},
					},
				},
				3: &models.Audience{
					ID:                3,
					Type:              models.AudienceType,
					Description:       "Sample Audience",
					Age:               25,
					AgeGroup:          "25-40",
					Gender:            "Male",
					BirthCountry:      "USA",
					HoursSpentOnMedia: 18,
					NumberOfPurchases: 4,
				},
			},
		},
		2: {
			ID: 2,
			Favourites: map[int]models.Asset{
				1: &models.Insight{
					ID:          1,
					Type:        models.InsightType,
					Description: "Sample Insight",
					Text:        "Sample Insight Text",
				},
				2: &models.Chart{
					ID:          2,
					Type:        models.ChartType,
					Description: "Sample Chart",
					Title:       "Sample Chart Title",
					XAxesTitle:  "X-Axis",
					YAxesTitle:  "Y-Axis",
					DataPoints: []models.Point{
						{X: 10, Y: 10},
						{X: 20, Y: 20},
					},
				},
				3: &models.Audience{
					ID:                3,
					Type:              models.AudienceType,
					Description:       "Sample Audience",
					Age:               25,
					AgeGroup:          "25-40",
					Gender:            "Male",
					BirthCountry:      "USA",
					HoursSpentOnMedia: 18,
					NumberOfPurchases: 4,
				},
			},
		},
		3: {
			ID: 3,
			Favourites: map[int]models.Asset{
				1: &models.Insight{
					ID:          1,
					Type:        models.InsightType,
					Description: "Sample Insight",
					Text:        "Sample Insight Text",
				},
				2: &models.Chart{
					ID:          2,
					Type:        models.ChartType,
					Description: "Sample Chart",
					Title:       "Sample Chart Title",
					XAxesTitle:  "X-Axis",
					YAxesTitle:  "Y-Axis",
					DataPoints: []models.Point{
						{X: 10, Y: 10},
						{X: 20, Y: 20},
					},
				},
				3: &models.Audience{
					ID:                3,
					Type:              models.AudienceType,
					Description:       "Sample Audience",
					Age:               25,
					AgeGroup:          "25-40",
					Gender:            "Male",
					BirthCountry:      "USA",
					HoursSpentOnMedia: 18,
					NumberOfPurchases: 4,
				},
			},
		},
	}

	// Create a new HTTP request
	var req *http.Request
	var err error
	if tc.payload != nil {
		body, _ := json.Marshal(tc.payload)
		req, err = http.NewRequest(tc.method, tc.url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(tc.method, tc.url, nil)
	}
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder and assign the handler
	rr := httptest.NewRecorder()
	handler := NewUserHandler(repo)

	// Create a new router and assign the handler
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}/favorites", handler.GetUserFavorites).Methods("GET")
	r.HandleFunc("/users/{id}/favorites", handler.AddUserFavorite).Methods("POST")
	r.HandleFunc("/users/{id}/favorites/{assetID}", handler.DeleteUserFavorite).Methods("DELETE")
	r.HandleFunc("/users/{id}/favorites/{assetID}", handler.EditUserFavorite).Methods("PUT")

	// Serve the HTTP request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != tc.expectedStatus {
		t.Errorf("handler returned wrong status code, got: %v expected: %v", status, tc.expectedStatus)
	}

	// Check the response body
	if tc.expectedBody != "" && !strings.Contains(rr.Body.String(), tc.expectedBody) {
		t.Errorf("handler returned unexpected body, got: %v expected: %v", rr.Body.String(), tc.expectedBody)
	}
}

// TestHandlers tests the GetUserFavorites, AddUserFavorite, DeleteUserFavorite, and EditUserFavorite handlers
func TestHandlers(t *testing.T) {
	getUserFavoritesTests := []testCase{
		{
			name:           "ValidUserFavorites",
			method:         "GET",
			url:            "/users/1/favorites",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "UserNotFound",
			method:         "GET",
			url:            "/users/999999/favorites",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "User not found",
		},
	}

	t.Run("GetUserFavorites", func(t *testing.T) {
		for _, tc := range getUserFavoritesTests {
			t.Run(tc.name, func(t *testing.T) {
				runTestCase(t, tc)
			})
		}
	})

	addUserFavoriteTests := []testCase{
		{
			name:   "ValidAddUserFavoriteInsight",
			method: "POST",
			url:    "/users/1/favorites",
			payload: &models.Insight{
				ID:          100, // ID is set to 100 to avoid conflict with existing assets
				Type:        models.InsightType,
				Description: "Sample Insight for testing to add as favorite",
				Text:        "Testing Insight",
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   "{\"id\":100,\"type\":\"Insight\",\"description\":\"Sample Insight for testing to add as favorite\",\"text\":\"Testing Insight\"}",
		},
		{
			name:   "ValidAddUserFavoriteChart",
			method: "POST",
			url:    "/users/2/favorites",
			payload: &models.Chart{
				ID:          200, // ID is set to 101 to avoid conflict with existing assets
				Type:        models.ChartType,
				Description: "Sample Chart for testing to add as favorite",
				Title:       "Testing Chart",
				XAxesTitle:  "X-Axis",
				YAxesTitle:  "Y-Axis",
				DataPoints: []models.Point{
					{X: 10, Y: 10},
					{X: 20, Y: 20},
				},
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   "{\"id\":200,\"type\":\"Chart\",\"description\":\"Sample Chart for testing to add as favorite\",\"title\":\"Testing Chart\",\"xAxesTitle\":\"X-Axis\",\"yAxesTitle\":\"Y-Axis\",\"dataPoints\":[{\"X\":10,\"Y\":10},{\"X\":20,\"Y\":20}]}",
		},
		{
			name:   "ValidAddUserFavoriteAudience",
			method: "POST",
			url:    "/users/3/favorites",
			payload: &models.Audience{
				ID:                300, // ID is set to 102 to avoid conflict with existing assets
				Type:              models.AudienceType,
				Description:       "Sample Audience for testing to add as favorite",
				Age:               15,
				AgeGroup:          "10-20",
				Gender:            "Male",
				BirthCountry:      "USA",
				HoursSpentOnMedia: 26,
				NumberOfPurchases: 8,
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   "{\"id\":300,\"type\":\"Audience\",\"description\":\"Sample Audience for testing to add as favorite\",\"age\":15,\"ageGroup\":\"10-20\",\"gender\":\"Male\",\"birthCountry\":\"USA\",\"hoursSpentOnMedia\":26,\"numberOfPurchases\":8}",
		},
		{
			name:   "AddUserFavoriteAssetExists",
			method: "POST",
			url:    "/users/1/favorites",
			payload: &models.Insight{
				ID:          1, // ID is set to 1 to match an existing asset
				Type:        models.InsightType,
				Description: "Sample Insight for testing to add existing as favorite",
				Text:        "Testing Insight",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Asset already exists",
		},
		{
			name:   "AddUserFavoriteInvalidAssetType",
			method: "POST",
			url:    "/users/1/favorites",
			payload: &models.Insight{
				ID:          400,
				Type:        "InvalidType", // Invalid asset type
				Description: "Sample Insight for testing to add as favorite",
				Text:        "Testing Insight",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid asset type",
		},
		{
			name:   "AddUserFavoriteUserNotFound",
			method: "POST",
			url:    "/users/999999/favorites",
			payload: &models.Insight{
				ID:          500,
				Type:        models.InsightType,
				Description: "Sample Insight for testing to add as favorite",
				Text:        "Testing Insight",
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "User not found",
		},
	}

	t.Run("AddUserFavorite", func(t *testing.T) {
		for _, tc := range addUserFavoriteTests {
			t.Run(tc.name, func(t *testing.T) {
				runTestCase(t, tc)
			})
		}
	})

	deleteUserFavoriteTests := []testCase{
		{
			name:           "ValidDeleteUserFavorite",
			method:         "DELETE",
			url:            "/users/1/favorites/1",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "DeleteUserFavoriteAssetNotFound",
			method:         "DELETE",
			url:            "/users/2/favorites/999999",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Asset not found",
		},
		{
			name:           "DeleteUserFavoriteUserNotFound",
			method:         "DELETE",
			url:            "/users/999999/favorites/1",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "User not found",
		},
	}

	t.Run("DeleteUserFavorite", func(t *testing.T) {
		for _, tc := range deleteUserFavoriteTests {
			t.Run(tc.name, func(t *testing.T) {
				runTestCase(t, tc)
			})
		}
	})

	editUserFavoriteTests := []testCase{
		{
			name:   "ValidEditUserFavoriteInsight",
			method: "PUT",
			url:    "/users/1/favorites/1",
			payload: &models.Insight{
				ID:          1,
				Type:        models.InsightType,
				Description: "Updated Insight",
				Text:        "Updated Insight Text",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "{\"id\":1,\"type\":\"Insight\",\"description\":\"Updated Insight\",\"text\":\"Updated Insight Text\"}",
		},
		{
			name:   "ValidEditUserFavoriteChart",
			method: "PUT",
			url:    "/users/2/favorites/2",
			payload: &models.Chart{
				ID:          2,
				Type:        models.ChartType,
				Description: "Updated Chart",
				Title:       "Updated Chart Title",
				XAxesTitle:  "X-Updated",
				YAxesTitle:  "Y-Updated",
				DataPoints: []models.Point{
					{X: 10, Y: 10},
					{X: 20, Y: 20},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "{\"id\":2,\"type\":\"Chart\",\"description\":\"Updated Chart\",\"title\":\"Updated Chart Title\",\"xAxesTitle\":\"X-Updated\",\"yAxesTitle\":\"Y-Updated\",\"dataPoints\":[{\"X\":10,\"Y\":10},{\"X\":20,\"Y\":20}]}",
		},
		{
			name:   "ValidEditUserFavoriteAudience",
			method: "PUT",
			url:    "/users/3/favorites/3",
			payload: &models.Audience{
				ID:                3,
				Type:              models.AudienceType,
				Description:       "Updated Audience",
				Age:               15,
				AgeGroup:          "10-20",
				Gender:            "Male",
				BirthCountry:      "USA",
				HoursSpentOnMedia: 40, // update the hours spent on media
				NumberOfPurchases: 20, // update the number of purchases
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "{\"id\":3,\"type\":\"Audience\",\"description\":\"Updated Audience\",\"age\":15,\"ageGroup\":\"10-20\",\"gender\":\"Male\",\"birthCountry\":\"USA\",\"hoursSpentOnMedia\":40,\"numberOfPurchases\":20}",
		},
		{
			name:   "EditUserFavoriteNoIDMatch",
			method: "PUT",
			url:    "/users/1/favorites/1",
			payload: &models.Insight{
				ID:          2, // ID does not match the asset ID in the URL (asset with ID 2 exists in the user's favorites)
				Type:        models.InsightType,
				Description: "Updated Insight",
				Text:        "Updated Insight Text",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Asset ID in the request body (ID: 2) does not match the asset ID in the URL (ID: 1)",
		},
		{
			name:   "EditUserFavoriteNoTypeMatch",
			method: "PUT",
			url:    "/users/3/favorites/3",
			payload: &models.Chart{ // the existing asset is Audience, not Chart
				ID:          3,
				Type:        models.ChartType,
				Description: "Updated Chart",
				Title:       "Updated Chart",
				XAxesTitle:  "X-Updated",
				YAxesTitle:  "Y-Updated",
				DataPoints: []models.Point{
					{X: 10, Y: 10},
					{X: 20, Y: 20},
				},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Asset type in the request body (Chart) does not match the existing asset type (Audience)",
		},
		{
			name:   "EditUserFavoriteUserNotFound",
			method: "PUT",
			url:    "/users/999999/favorites/1",
			payload: &models.Insight{
				ID:          1,
				Type:        models.InsightType,
				Description: "Updated Insight",
				Text:        "Updated Insight Text",
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "User not found",
		},
		{
			name:   "EditUserFavoriteAssetNotFound",
			method: "PUT",
			url:    "/users/1/favorites/999999",
			payload: &models.Insight{
				ID:          999999, // Using the ID that does not exist in the favorites
				Type:        models.InsightType,
				Description: "Updated Insight",
				Text:        "Updated Insight Text",
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Asset not found",
		},
		{
			name:   "EditUserFavoriteInvalidAssetType",
			method: "PUT",
			url:    "/users/1/favorites/1",
			payload: &models.Insight{
				ID:          1,
				Type:        "InvalidType", // Invalid asset type
				Description: "Updated Insight",
				Text:        "Updated Insight Text",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid asset type",
		},
	}

	t.Run("EditUserFavorite", func(t *testing.T) {
		for _, tc := range editUserFavoriteTests {
			t.Run(tc.name, func(t *testing.T) {
				runTestCase(t, tc)
			})
		}
	})
}
