package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/swaggo/swag"
	"gitlab.com/ulexxander/remoconf/docs"
	"gitlab.com/ulexxander/remoconf/restapi"
	"gitlab.com/ulexxander/remoconf/service/projects"
	"gitlab.com/ulexxander/remoconf/service/users"
	"gitlab.com/ulexxander/remoconf/storage/sqlite"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	if err := run(logger); err != nil {
		logger.Fatalf("fatal error: %s", err)
	}
}

func run(logger *log.Logger) error {
	logger.Println("TDD FTW, we dont care for main for now")

	db, err := sqlx.Connect("sqlite3", ".temp/sqlite.db")
	if err != nil {
		return errors.Wrap(err, "error connecting to sqlite")
	}

	if err := sqlite.Migrate(db); err != nil {
		return errors.Wrap(err, "migrating sqlite")
	}

	users := users.NewService(sqlite.NewUsersStore(db))
	projects := projects.NewService(sqlite.NewProjectsStore(db))

	docs.SwaggerInfo.Title = "Remoconf API"
	swaggerDocs, err := swag.ReadDoc()
	if err != nil {
		return errors.Wrap(err, "reading swagger")
	}

	h := restapi.NewHandler(users, projects, []byte(swaggerDocs), logger)

	port := ":4000"
	logger.Println("listening on", port)
	err = http.ListenAndServe(port, h)
	if err != nil {
		return errors.Wrap(err, "listening http")
	}

	return nil
}
