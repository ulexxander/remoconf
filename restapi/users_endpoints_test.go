package restapi_test

import (
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/ulexxander/remoconf/restapi"
	"gitlab.com/ulexxander/remoconf/storage"
	"gitlab.com/ulexxander/remoconf/storage/sqlite"
	"gitlab.com/ulexxander/remoconf/testutil"
)

func TestUsersEndpoints(t *testing.T) {
	client := setupAPIClient(t)

	login := "alex"
	password := "123"

	var userID int

	t.Run("creating user", func(t *testing.T) {
		var resBody struct {
			Data struct {
				ID int
			}
			restapi.ResponseError
		}
		res := client.Post(t, "/users", storage.UserCreateParams{
			Login:    login,
			Password: password,
		}, &resBody)
		require.Empty(t, resBody.Error)
		require.Equal(t, 200, res.StatusCode)
		userID = resBody.Data.ID
	})

	t.Run("getting created user", func(t *testing.T) {
		var resBody struct {
			Data storage.User
			restapi.ResponseError
		}
		res := client.Get(t, fmt.Sprintf("/users/%d", userID), &resBody)
		require.Empty(t, resBody.Error)
		require.Equal(t, 200, res.StatusCode)

		require.Equal(t, userID, resBody.Data.ID)
		require.Equal(t, login, resBody.Data.Login)
		// TODO: hash
		require.Equal(t, password, resBody.Data.Password)
		require.NotZero(t, resBody.Data.CreatedAt)
		require.Nil(t, resBody.Data.UpdatedAt)
	})
}

func setupAPIClient(t *testing.T) *testutil.APIClient {
	db := testutil.SetupDBTest(t)
	us := sqlite.NewUsersStore(db)
	logger := log.New(io.Discard, "", log.LstdFlags)
	h := restapi.NewHandler(us, logger)
	serv := httptest.NewServer(h)
	t.Cleanup(func() {
		serv.Close()
	})
	return testutil.NewAPIClient(serv)
}
