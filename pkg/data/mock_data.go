// data/mock_data.go

package data

import (
	"fmt"

	"github.com/ceciivanov/go-challenge/pkg/models"
	"github.com/ceciivanov/go-challenge/pkg/utils"
)

// Define a global variable map to store users
var Users map[int]models.User

// Initialize the Users map when the package is loaded
func init() {
	Users = make(map[int]models.User)
}

func GenerateMockData(numberOfUsers, numberOfAssets int) {
	for i := 1; i <= numberOfUsers; i++ {
		userID := i
		user := models.User{
			ID:         userID,
			Favourites: make(map[int]models.Asset),
		}

		for j := 1; j <= numberOfAssets; j++ {
			assetID := j
			assetType := j % 3

			var asset models.Asset
			switch assetType {
			case 0:
				asset = &models.Chart{
					ID:          assetID,
					Type:        models.AssetTypeChart,
					Description: "Sample Chart for GWI",
					Title:       fmt.Sprintf("GWI Chart %d", j),
					XAxesTitle:  "X-Axis",
					YAxesTitle:  "Y-Axis",
					DataPoints:  utils.GetRandomPoints(1, 5),
				}
			case 1:
				asset = &models.Insight{
					ID:          assetID,
					Type:        models.AssetTypeInsight,
					Description: "Sample Insight for GWI",
					Text:        fmt.Sprintf("GWI Insight %d", j),
				}
			case 2:
				asset = &models.Audience{
					ID:                assetID,
					Type:              models.AssetTypeAudience,
					Description:       "Sample Audience for GWI",
					Age:               utils.GetRandomNumber(100),
					AgeGroup:          utils.GetRandomAgeGroup(),
					Gender:            utils.GetRandomGender(),
					BirthCountry:      utils.GetRandomCountry(),
					HoursSpentOnMedia: utils.GetRandomNumber(100),
					NumberOfPurchases: utils.GetRandomNumber(100),
				}
			}

			user.Favourites[assetID] = asset
		}

		Users[userID] = user
	}
}
