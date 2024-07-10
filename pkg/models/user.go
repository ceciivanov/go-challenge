package models

type User struct {
	ID         string           `json:"id"`
	Favourites map[string]Asset `json:"favourites"`
}
