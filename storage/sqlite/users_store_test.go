package sqlite_test

import (
	"testing"

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

	login := "alex"
	password := "123"

	var userID int

	t.Run("creating user", func(t *testing.T) {
		id, err := us.Create(storage.UserCreateParams{
			Login:    login,
			Password: password,
		})
		require.NoError(t, err)
		userID = id
	})

	t.Run("getting created user", func(t *testing.T) {
		u, err := us.GetByID(userID)
		require.NoError(t, err)
		require.NotNil(t, u)

		require.Equal(t, userID, u.ID)
		require.Equal(t, login, u.Login)
		require.Equal(t, password, u.Password)
		require.NotZero(t, u.CreatedAt)
		require.Nil(t, u.UpdatedAt)
	})
}
