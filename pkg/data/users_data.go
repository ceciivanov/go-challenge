package data

import (
	"github.com/ceci/go-challenge/pkg/models"
)

var Users map[string]models.User

func init() {
	Users = make(map[string]models.User)
}
