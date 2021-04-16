package controller

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rikakobayashi/go-mvc-architecture/pkg/model"
	"github.com/rikakobayashi/go-mvc-architecture/pkg/presenter"
	"github.com/rikakobayashi/go-mvc-architecture/pkg/view"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("../db/database.sqlite3"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	err := db.Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			presenter.ErrorNotFound(w, err)
			return
		} else {
			presenter.Error(w, err)
			return
		}
	}
	presenter.Response(w, view.Users(users))
}

func User(w http.ResponseWriter, r *http.Request) {
	var user model.User
	id := chi.URLParam(r, "id")

	err := db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			presenter.ErrorNotFound(w, err)
			return
		} else {
			presenter.Error(w, err)
			return
		}
	}

	u := view.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	presenter.Response(w, u)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
	}

	err := db.Create(&user).Error
	if err != nil {
		presenter.Error(w, err)
		return
	}
	presenter.Response(w, user)
}
