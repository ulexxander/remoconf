package restapi

import (
	"log"

	"github.com/go-chi/chi/v5"
	"gitlab.com/ulexxander/remoconf/storage"
)

type Handler struct {
	*chi.Mux

	logger     *log.Logger
	usersStore storage.UsersStore
}

func NewHandler(us storage.UsersStore, logger *log.Logger) *Handler {
	mux := chi.NewMux()

	h := Handler{
		Mux:        mux,
		logger:     logger,
		usersStore: us,
	}

	mux.Get("/users/{id}", h.getUserByID)
	mux.Post("/users", h.postUser)

	return &h
}
