package testutil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gitlab.com/ulexxander/remoconf/restapi"
	"gitlab.com/ulexxander/remoconf/storage"
)

type APIClient struct {
	serv *httptest.Server
}

func NewAPIClient(serv *httptest.Server) *APIClient {
	return &APIClient{serv: serv}
}

func (ac *APIClient) Request(t *testing.T, method, endpoint string, reqBody, resBody interface{}) *http.Response {
	res, err := ac.doRequest(method, endpoint, reqBody, resBody)
	if err != nil {
		t.Fatalf("failed doing %s %s: %s", method, endpoint, err)
	}
	return res
}

func (ac *APIClient) doRequest(method, endpoint string, reqBody, resBody interface{}) (*http.Response, error) {
	var bodyReader io.Reader
	if reqBody != nil {
		encoded, err := json.Marshal(reqBody)
		if err != nil {
			return nil, errors.Wrap(err, "marshaling request body")
		}
		bodyReader = bytes.NewBuffer(encoded)
	}

	req, err := http.NewRequest(method, ac.serv.URL+endpoint, bodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "initializing request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "doing request")
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading response body")
	}

	err = json.Unmarshal(b, resBody)
	if err != nil {
		content := b
		if len(content) > 100 {
			content = content[:100]
		}
		return nil, errors.Wrapf(err, "unmarshaling body, first 100 bytes: %s", content)
	}

	return res, nil
}

func (ac *APIClient) Get(t *testing.T, endpoint string, resBody interface{}) *http.Response {
	return ac.Request(t, "GET", endpoint, nil, resBody)
}

func (ac *APIClient) Post(t *testing.T, endpoint string, reqBody, resBody interface{}) *http.Response {
	return ac.Request(t, "POST", endpoint, reqBody, resBody)
}

func (ac *APIClient) CreateUser(t *testing.T, login, password string) *storage.CreatedItem {
	var resBody struct {
		Data storage.CreatedItem
		restapi.ResponseError
	}
	res := ac.Post(t, "/users", storage.UserCreateParams{
		Login:    login,
		Password: password,
	}, &resBody)
	require.Empty(t, resBody.Error)
	require.Equal(t, 200, res.StatusCode)
	return &resBody.Data
}

func (ac *APIClient) CreateUserDefault(t *testing.T) *storage.CreatedItem {
	return ac.CreateUser(t, "alex", "123")
}
