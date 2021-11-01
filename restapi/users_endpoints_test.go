package restapi_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/ulexxander/remoconf/restapi"
	"gitlab.com/ulexxander/remoconf/storage"
	"gitlab.com/ulexxander/remoconf/testutil"
)

func TestUsersEndpoints(t *testing.T) {
	client := testutil.SetupRestAPI(t)

	login := "alex"
	password := "123"

	var userID int

	t.Run("creating user", func(t *testing.T) {
		created := client.CreateUser(t, login, password)
		userID = created.ID
	})

	t.Run("getting created user", func(t *testing.T) {
		var resBody struct {
			Data *storage.User
			restapi.ResponseError
		}
		res := client.Get(t, fmt.Sprintf("/users/%d", userID), &resBody)
		require.Empty(t, resBody.Error)
		require.Equal(t, 200, res.StatusCode)

		require.Equal(t, userID, resBody.Data.ID)
		require.Equal(t, login, resBody.Data.Login)
		require.NotEqual(t, password, resBody.Data.Password, "password is hashed")
		require.NotZero(t, resBody.Data.CreatedAt)
		require.Nil(t, resBody.Data.UpdatedAt)
	})
}
