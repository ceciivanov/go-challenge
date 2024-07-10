// data/mock_data.go

package data

import (
	"fmt"

	"github.com/ceciivanov/go-challenge/pkg/models"
)

func GenerateMockData() {
	for i := 1; i <= 5; i++ {
		userID := fmt.Sprintf("%d", i)
		user := models.User{
			ID:         userID,
			Favourites: make(map[string]models.Asset),
		}

		// Generate 100 assets for each user
		for j := 1; j <= 3; j++ {
			assetID := fmt.Sprintf("%d", j)
			assetType := j % 3

			var asset models.Asset
			switch assetType {
			case 0:
				asset = &models.Chart{
					ID:          assetID,
					Type:        models.AssetTypeChart,
					Description: "A sample chart",
					Title:       "Sample Chart",
					XAxesTitle:  "X-Axis",
					YAxesTitle:  "Y-Axis",
					DataPoints:  []models.Point{{X: 1, Y: 2}, {X: 2, Y: 3}},
				}
			case 1:
				asset = &models.Insight{
					ID:          assetID,
					Type:        models.AssetTypeInsight,
					Description: "A sample insight",
					Text:        "Sample text",
				}
			case 2:
				asset = &models.Audience{
					ID:                assetID,
					Type:              models.AssetTypeAudience,
					Description:       "A sample audience",
					Age:               25,
					AgeGroup:          "24-35",
					Gender:            "Male",
					BirthCountry:      "Country",
					HoursSpentOnMedia: 3,
					NumberOfPurchases: 5,
				}
			}

			user.Favourites[assetID] = asset
		}

		Users[userID] = user
	}
}
