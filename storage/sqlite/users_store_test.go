package sqlite_test

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"gitlab.com/ulexxander/remoconf/storage"
	"gitlab.com/ulexxander/remoconf/storage/sqlite"
	"gitlab.com/ulexxander/remoconf/testutil"
)

func TestUsersStore(t *testing.T) {
	db := testutil.SetupDBTest(t)
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
