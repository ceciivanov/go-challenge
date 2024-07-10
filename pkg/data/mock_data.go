// data/mock_data.go

package data

import (
	"fmt"
	"math/rand"

	"github.com/ceciivanov/go-challenge/pkg/models"
	"github.com/ceciivanov/go-challenge/pkg/utils"
)

// Age groups
const (
	AgeGroupTeen   string = "0-17"
	AgeGroupAdult  string = "18-64"
	AgeGroupSenior string = "65+"
)

// Genders
const (
	Male      string = "Male"
	Female    string = "Female"
	NonBinary string = "Non-Binary"
)

var countries = []string{
	"Afghanistan", "Albania", "Algeria", "Andorra", "Angola", "Argentina", "Armenia", "Australia", "Austria", "Azerbaijan",
	"Bahamas", "Bahrain", "Bangladesh", "Barbados", "Belarus", "Belgium", "Belize", "Benin", "Bhutan", "Bolivia",
	"Bosnia and Herzegovina", "Botswana", "Brazil", "Brunei", "Bulgaria", "Burkina Faso", "Burundi", "Cambodia", "Cameroon", "Canada",
	"Cape Verde", "Central African Republic", "Chad", "Chile", "China", "Colombia", "Comoros", "Congo", "Costa Rica", "Croatia",
	"Cuba", "Cyprus", "Czech Republic", "Denmark", "Djibouti", "Dominica", "Dominican Republic", "Ecuador", "Egypt", "El Salvador",
	"Equatorial Guinea", "Eritrea", "Estonia", "Eswatini", "Ethiopia", "Fiji", "Finland", "France", "Gabon", "Gambia",
	"Georgia", "Germany", "Ghana", "Greece", "Grenada", "Guatemala", "Guinea", "Guinea-Bissau", "Guyana", "Haiti",
	"Honduras", "Hungary", "Iceland", "India", "Indonesia", "Iran", "Iraq", "Ireland", "Israel", "Italy",
	"Jamaica", "Japan", "Jordan", "Kazakhstan", "Kenya", "Kiribati", "Kuwait", "Kyrgyzstan", "Laos", "Latvia",
	"Lebanon", "Lesotho", "Liberia", "Libya", "Liechtenstein", "Lithuania", "Luxembourg", "Madagascar", "Malawi", "Malaysia",
	"Maldives", "Mali", "Malta", "Marshall Islands", "Mauritania", "Mauritius", "Mexico", "Micronesia", "Moldova", "Monaco",
	"Mongolia", "Montenegro", "Morocco", "Mozambique", "Myanmar", "Namibia", "Nauru", "Nepal", "Netherlands", "New Zealand",
	"Nicaragua", "Niger", "Nigeria", "North Macedonia", "Norway", "Oman", "Pakistan", "Palau", "Panama", "Papua New Guinea",
	"Paraguay", "Peru", "Philippines", "Poland", "Portugal", "Qatar", "Romania", "Russia", "Rwanda", "Saint Kitts and Nevis",
	"Saint Lucia", "Saint Vincent and the Grenadines", "Samoa", "San Marino", "Sao Tome and Principe", "Saudi Arabia", "Senegal", "Serbia", "Seychelles", "Sierra Leone",
	"Singapore", "Slovakia", "Slovenia", "Solomon Islands", "Somalia", "South Africa", "South Sudan", "Spain", "Sri Lanka", "Sudan",
	"Suriname", "Sweden", "Switzerland", "Syria", "Taiwan", "Tajikistan", "Tanzania", "Thailand", "Timor-Leste", "Togo",
	"Tonga", "Trinidad and Tobago", "Tunisia", "Turkey", "Turkmenistan", "Tuvalu", "Uganda", "Ukraine", "United Arab Emirates", "United Kingdom",
	"United States", "Uruguay", "Uzbekistan", "Vanuatu", "Vatican City", "Venezuela", "Vietnam", "Yemen", "Zambia", "Zimbabwe",
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
					DataPoints:  utils.GenerateRandomPoints(1, 5),
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
					Age:               uint(rand.Int()),
					AgeGroup:          utils.GetRandomAgeGroup(),
					Gender:            utils.GetRandomGender(),
					BirthCountry:      utils.GetRandomCountry(),
					HoursSpentOnMedia: uint(rand.Int()),
					NumberOfPurchases: uint(rand.Int()),
				}
			}

			user.Favourites[assetID] = asset
		}

		Users[userID] = user
	}
}
