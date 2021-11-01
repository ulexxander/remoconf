package testutil

import (
	"io"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"gitlab.com/ulexxander/remoconf/restapi"
	"gitlab.com/ulexxander/remoconf/service/projects"
	"gitlab.com/ulexxander/remoconf/service/users"
	"gitlab.com/ulexxander/remoconf/storage/sqlite"
)

func SetupRestAPI(t *testing.T) *APIClient {
	db := SetupDBTest(t)
	users := users.NewService(sqlite.NewUsersStore(db))
	projects := projects.NewService(sqlite.NewProjectsStore(db))
	logger := log.New(io.Discard, "", log.LstdFlags)

	e := restapi.Endpoints{
		Users:       users,
		Projects:    projects,
		SwaggerDocs: nil,
		Logger:      logger,
	}

	mux := chi.NewMux()
	e.Register(mux)

	serv := httptest.NewServer(mux)
	t.Cleanup(func() {
		serv.Close()
	})

	return NewAPIClient(serv)
}
