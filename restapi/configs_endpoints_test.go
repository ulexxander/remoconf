package restapi_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/ulexxander/remoconf/restapi"
	"gitlab.com/ulexxander/remoconf/storage"
	"gitlab.com/ulexxander/remoconf/testutil"
)

func TestConfigsEndpoints(t *testing.T) {
	client := testutil.SetupRestAPI(t)

	user := client.CreateUserDefault(t)
	project := client.CreateProject(t, "proj", "desc", user.ID)

	version := 1
	content := "{super cfg}"

	var configID int

	t.Run("creating config", func(t *testing.T) {
		var resBody struct {
			Data *storage.CreatedItem
			restapi.ResponseError
		}
		res := client.Post(t, "/configs", storage.ConfigCreateParams{
			ProjectID: project.ID,
			Version:   version,
			Content:   content,
			CreatedBy: user.ID,
		}, &resBody)
		require.Equal(t, 200, res.StatusCode)
		require.Empty(t, resBody.Error)
		configID = resBody.Data.ID
	})

	t.Run("created config listed", func(t *testing.T) {
		var resBody struct {
			Data []storage.Config
			restapi.ResponseError
		}
		res := client.Get(t, "/configs", &resBody)
		require.Equal(t, 200, res.StatusCode)
		require.Empty(t, resBody.Error)
		require.Len(t, resBody.Data, 1)
		require.Equal(t, configID, resBody.Data[0].ID)
	})
}
