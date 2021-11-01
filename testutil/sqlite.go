package testutil

import (
	"path/filepath"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"gitlab.com/ulexxander/remoconf/storage/sqlite"
)

func SetupDBTest(t *testing.T) *sqlx.DB {
	db, err := SetupDB(t.TempDir())
	if err != nil {
		t.Fatalf("failed to setup db: %s", err)
	}
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Logf("error closing db: %s", err)
		}
	})
	return db
}

func SetupDB(dir string) (*sqlx.DB, error) {
	path := filepath.Join(dir, "sqlite.db")
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, errors.Wrap(err, "opening")
	}
	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "pinging")
	}
	if err := sqlite.Migrate(db); err != nil {
		return nil, errors.Wrap(err, "migrating")
	}
	return db, nil
}
