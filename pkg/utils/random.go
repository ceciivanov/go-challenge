package utils

import (
	"math/rand"

	"github.com/ceciivanov/go-challenge/pkg/data/"
	"github.com/ceciivanov/go-challenge/pkg/models"
)

// GetRandomAgeGroup returns a random age group
func GetRandomAgeGroup() string {
	ageGroups := []string{data.AgeGroupTeen, data.AgeGroupAdult, data.AgeGroupSenior}
	return ageGroups[rand.Intn(len(ageGroups))]
}

// GetRandomGender returns a random gender
func GetRandomGender() string {
	genders := []string{data.Male, data.Female, data.NonBinary}
	return genders[rand.Intn(len(genders))]
}

// GenerateRandomPoints generates a random number of random points
func GenerateRandomPoints(minPoints, maxPoints int) []models.Point {
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

// GetRandomCountry returns a random country from the predefined list
func GetRandomCountry() string {
	return data.countries[rand.Intn(len(data.countries))]
}
