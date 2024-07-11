package mock_data

import (
	"fmt"

	"github.com/ceciivanov/go-challenge/pkg/models"
)

// GenerateMockData generates mock data for users and assets and returns a map of users
func GenerateMockData(NumberOfUsers, NumberOfAssets int) map[int]models.User {
	Users := make(map[int]models.User)

	for i := 1; i <= NumberOfUsers; i++ {
		userID := i
		user := models.User{
			ID:         userID,
			Favourites: make(map[int]models.Asset),
		}

		for j := 1; j <= NumberOfAssets; j++ {
			assetID := j
			assetType := j % 3

			var asset models.Asset
			switch assetType {
			case 0:
				asset = &models.Chart{
					ID:          assetID,
					Type:        models.ChartType,
					Description: "Sample Chart for GWI",
					Title:       fmt.Sprintf("GWI Chart %d", j),
					XAxesTitle:  "X-Axis",
					YAxesTitle:  "Y-Axis",
					DataPoints:  GetRandomPoints(1, 5),
				}
			case 1:
				asset = &models.Insight{
					ID:          assetID,
					Type:        models.InsightType,
					Description: "Sample Insight for GWI",
					Text:        fmt.Sprintf("GWI Insight %d", j),
				}
			case 2:
				asset = &models.Audience{
					ID:                assetID,
					Type:              models.AudienceType,
					Description:       "Sample Audience for GWI",
					Age:               GetRandomNumber(100),
					AgeGroup:          GetRandomAgeGroup(),
					Gender:            GetRandomGender(),
					BirthCountry:      GetRandomCountry(),
					HoursSpentOnMedia: GetRandomNumber(100),
					NumberOfPurchases: GetRandomNumber(100),
				}
			}

			user.Favourites[assetID] = asset
		}

		Users[userID] = user
	}

	return Users
}
