package utils

import (
	"math/rand"

	"github.com/ceciivanov/go-challenge/pkg/models"
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

// GetRandomPoints generates a random number of random points
func GetRandomPoints(minPoints, maxPoints int) []models.Point {
	numPoints := rand.Intn(maxPoints-minPoints+1) + minPoints
	points := make([]models.Point, numPoints)
	for i := 0; i < numPoints; i++ {
		points[i] = models.Point{
			X: rand.Float32(), // Adjust range as needed
			Y: rand.Float32(), // Adjust range as needed
		}
	}
	return points
}

// GetRandomAgeGroup returns a random age group
func GetRandomAgeGroup() string {
	ageGroups := []string{AgeGroupTeen, AgeGroupAdult, AgeGroupSenior}
	return ageGroups[rand.Intn(len(ageGroups))]
}

// GetRandomGender returns a random gender
func GetRandomGender() string {
	genders := []string{Male, Female, NonBinary}
	return genders[rand.Intn(len(genders))]
}

// GetRandomCountry returns a random country from the predefined list
func GetRandomCountry() string {
	return countries[rand.Intn(len(countries))]
}

func GetRandomNumber(max int) uint {
	return uint(rand.Intn(max))
}
