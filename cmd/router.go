package main

import (
	"github.com/rikakobayashi/go-mvc-architecture/pkg/controller"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/users", controller.UserIndex)
	r.Get("/user/{id}", controller.User)
	r.Post("/user/create", controller.CreateUser)
	return r
}
