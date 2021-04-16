package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rikakobayashi/go-mvc-architecture/pkg/model"
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
	db.Find(&users)
	u := view.Users(&users)
	e := json.NewEncoder(w)
	if err := e.Encode(u); err != nil {
		fmt.Println(err)
		return
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	var user model.User
	id := chi.URLParam(r, "id")
	db.Find(&user, id)
	u := view.User{
		Id: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
	e := json.NewEncoder(w)
	if err := e.Encode(u); err != nil {
		fmt.Println(err)
		return
	}
}

func UserNew(w http.ResponseWriter, r *http.Request) {
	user := model.User{
		Name: r.FormValue("name"),
		Email: r.FormValue("email"),
	}
	result := db.Create(&user)
	
	err := result.Error
	if err != nil {
		fmt.Println(err)
		return
	}
}