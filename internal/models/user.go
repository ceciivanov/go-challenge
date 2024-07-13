package models

type User struct {
	ID         int           `json:"id"`
	Favourites map[int]Asset `json:"favourites"`
}
