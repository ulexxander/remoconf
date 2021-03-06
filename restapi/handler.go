package restapi

import (
	"log"

	"github.com/go-chi/chi/v5"
	"gitlab.com/ulexxander/remoconf/service/configs"
	"gitlab.com/ulexxander/remoconf/service/projects"
	"gitlab.com/ulexxander/remoconf/service/users"
)

type Endpoints struct {
	Users       *users.Service
	Projects    *projects.Service
	Configs     *configs.Service
	SwaggerDocs []byte
	Logger      *log.Logger
}

func (e *Endpoints) Register(m *chi.Mux) {
	m.Get("/users/{id}", e.GetUserByID)
	m.Post("/users", e.PostUser)

	m.Get("/projects", e.GetProjectsAll)
	m.Post("/projects", e.PostProject)

	m.Get("/projects/{id}/configs", e.GetConfigsByProject)
	m.Post("/configs", e.PostConfig)

	m.Get("/swagger/docs.json", e.GetSwaggerDocs)
	m.Get("/swagger/*", e.GetSwaggerWebInterface())
}
