package restapi_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/ulexxander/remoconf/restapi"
	"gitlab.com/ulexxander/remoconf/storage"
	"gitlab.com/ulexxander/remoconf/testutil"
)

func TestProjectsEndpoints(t *testing.T) {
	client := testutil.SetupRestAPI(t)

	user := client.CreateUserDefault(t)
	title := "super proj"
	description := "abc"

	var projectID int

	t.Run("creating project", func(t *testing.T) {
		var resBody struct {
			Data *storage.CreatedItem
			restapi.ResponseError
		}
		res := client.Post(t, "/projects", storage.ProjectCreateParams{
			Title:       title,
			Description: description,
			CreatedBy:   user.ID,
		}, &resBody)
		require.Equal(t, 200, res.StatusCode)
		require.Empty(t, resBody.Error)
		projectID = resBody.Data.ID
	})

	t.Run("created project listed", func(t *testing.T) {
		var resBody struct {
			Data []storage.Project
			restapi.ResponseError
		}
		res := client.Get(t, "/projects", &resBody)
		require.Equal(t, 200, res.StatusCode)
		require.Empty(t, resBody.Error)
		require.Len(t, resBody.Data, 1)
		require.Equal(t, projectID, resBody.Data[0].ID)
	})
}
