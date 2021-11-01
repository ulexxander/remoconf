package restapi

import (
	"log"

	"github.com/go-chi/chi/v5"
	"gitlab.com/ulexxander/remoconf/service/users"
)

type Handler struct {
	*chi.Mux

	logger *log.Logger
	users  *users.Service
}

func NewHandler(users *users.Service, logger *log.Logger) *Handler {
	mux := chi.NewMux()

	h := Handler{
		Mux:    mux,
		logger: logger,
		users:  users,
	}

	mux.Get("/users/{id}", h.getUserByID)
	mux.Post("/users", h.postUser)

	return &h
}
