package restapi

import (
	"log"

	"github.com/go-chi/chi/v5"
	"gitlab.com/ulexxander/remoconf/service/projects"
	"gitlab.com/ulexxander/remoconf/service/users"
)

type Handler struct {
	*chi.Mux

	users       *users.Service
	projects    *projects.Service
	swaggerDocs []byte

	logger *log.Logger
}

func NewHandler(users *users.Service, projects *projects.Service, swaggerDocs []byte, logger *log.Logger) *Handler {
	mux := chi.NewMux()

	h := Handler{
		Mux:         mux,
		users:       users,
		projects:    projects,
		swaggerDocs: swaggerDocs,
		logger:      logger,
	}

	mux.Get("/swagger/docs.json", h.getSwaggerDocs)
	mux.Get("/swagger/*", h.getSwaggerWebInterface())

	mux.Get("/users/{id}", h.GetUserByID)
	mux.Post("/users", h.PostUser)

	mux.Get("/projects", h.getProjectsAll)
	mux.Post("/projects", h.PostProject)

	return &h
}
