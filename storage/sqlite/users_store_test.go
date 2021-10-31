package sqlite_test

import (
	"path/filepath"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gitlab.com/ulexxander/remoconf/storage"
	"gitlab.com/ulexxander/remoconf/storage/sqlite"
)

func TestUsersStore(t *testing.T) {
	db := setupDBTest(t)
	us := sqlite.NewUsersStore(db)

	t.Run("user does not exist", func(t *testing.T) {
		u, err := us.GetByID(1)
		require.Error(t, err)
		require.Nil(t, u)
	})

	var createdUserID int

	t.Run("creating user", func(t *testing.T) {
		id, err := us.Create(storage.UserCreateParams{
			Login:    "alex",
			Password: "123",
		})
		require.NoError(t, err)
		createdUserID = id
	})

	t.Run("getting created user", func(t *testing.T) {
		u, err := us.GetByID(createdUserID)
		require.NoError(t, err)
		require.NotNil(t, u)
	})
}

func setupDBTest(t *testing.T) *sqlx.DB {
	db, err := setupDB(t.TempDir())
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

func setupDB(dir string) (*sqlx.DB, error) {
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
