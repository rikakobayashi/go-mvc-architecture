package view

import (
	"github.com/rikakobayashi/go-mvc-architecture/pkg/model"
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Users(u []model.User) []User {
	users := make([]User, len(u))
	for i, e := range u {
		users[i] = User{
			ID:    e.ID,
			Name:  e.Name,
			Email: e.Email,
		}
	}
	return users
}
