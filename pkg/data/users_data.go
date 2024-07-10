package data

import (
	"github.com/ceciivanov/go-challenge/pkg/models"
)

var Users map[int]models.User

func init() {
	Users = make(map[int]models.User)
}
