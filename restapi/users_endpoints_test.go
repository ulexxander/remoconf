package restapi_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/ulexxander/remoconf/restapi"
	"gitlab.com/ulexxander/remoconf/storage"
	"gitlab.com/ulexxander/remoconf/storage/sqlite"
	"gitlab.com/ulexxander/remoconf/testutil"
)

func TestUsersEndpoints(t *testing.T) {
	client := setupTestClient(t)

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

type TestClient struct {
	serv *httptest.Server
}

func NewTestClient(serv *httptest.Server) *TestClient {
	return &TestClient{serv: serv}
}

func (tc *TestClient) Request(t *testing.T, method, endpoint string, reqBody, resBody interface{}) *http.Response {
	var bodyReader io.Reader
	if reqBody != nil {
		encoded, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatalf("could not json marshal body: %s", err)
		}
		bodyReader = bytes.NewBuffer(encoded)
	}

	req, err := http.NewRequest(method, tc.serv.URL+endpoint, bodyReader)
	if err != nil {
		t.Fatalf("failed to prepare new request: %s", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("error when doing %s request to %s: %s", method, endpoint, err)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("error when reading body: %s", err)
	}

	err = json.Unmarshal(b, resBody)
	if err != nil {
		content := b
		if len(content) > 100 {
			content = content[:100]
		}
		t.Fatalf("failed to unmarshal body: %s, first 100 bytes: %s", err, content)
	}

	return res
}

func (tc *TestClient) Get(t *testing.T, endpoint string, resBody interface{}) *http.Response {
	return tc.Request(t, "GET", endpoint, nil, resBody)
}

func (tc *TestClient) Post(t *testing.T, endpoint string, reqBody, resBody interface{}) *http.Response {
	return tc.Request(t, "POST", endpoint, reqBody, resBody)
}

func setupTestClient(t *testing.T) *TestClient {
	db := testutil.SetupDBTest(t)
	us := sqlite.NewUsersStore(db)
	logger := log.New(io.Discard, "", log.LstdFlags)
	h := restapi.NewHandler(us, logger)
	serv := httptest.NewServer(h)
	t.Cleanup(func() {
		serv.Close()
	})
	return &TestClient{serv}
}
