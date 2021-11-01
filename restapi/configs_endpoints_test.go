package restapi_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/ulexxander/remoconf/restapi"
	"gitlab.com/ulexxander/remoconf/service/configs"
	"gitlab.com/ulexxander/remoconf/storage"
	"gitlab.com/ulexxander/remoconf/testutil"
)

func TestConfigsEndpoints_CRUD(t *testing.T) {
	client := testutil.SetupRestAPI(t)

	user := client.CreateUserDefault(t)
	project := client.CreateProject(t, "proj", "desc", user.ID)

	content := "{super cfg}"

	var configID int

	t.Run("creating config", func(t *testing.T) {
		var resBody struct {
			Data *storage.CreatedItem
			restapi.ResponseError
		}
		res := client.Post(t, "/configs", configs.ConfigCreateParams{
			ProjectID: project.ID,
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
		res := client.Get(t, fmt.Sprintf("/projects/%d/configs", project.ID), &resBody)
		require.Equal(t, 200, res.StatusCode)
		require.Empty(t, resBody.Error)
		require.Len(t, resBody.Data, 1)
		require.Equal(t, configID, resBody.Data[0].ID)
	})
}

func TestConfigsEndpoints_IncreasesVersion(t *testing.T) {
	client := testutil.SetupRestAPI(t)

	user := client.CreateUserDefault(t)
	project := client.CreateProject(t, "proj2", "desc", user.ID)

	var resConf1 struct{ Data *storage.CreatedItem }
	client.Post(t, "/configs", configs.ConfigCreateParams{
		ProjectID: project.ID,
		CreatedBy: user.ID,
	}, &resConf1)

	var resConf2 struct{ Data *storage.CreatedItem }
	client.Post(t, "/configs", configs.ConfigCreateParams{
		ProjectID: project.ID,
		CreatedBy: user.ID,
	}, &resConf2)

	var resConfs struct{ Data []storage.Config }
	client.Get(t, fmt.Sprintf("/projects/%d/configs", project.ID), &resConfs)
	require.Len(t, resConfs.Data, 2)

	require.Equal(t, resConf1.Data.ID, resConfs.Data[0].ID)
	require.Equal(t, 1, resConfs.Data[0].Version)

	require.Equal(t, resConf2.Data.ID, resConfs.Data[1].ID)
	require.Equal(t, 2, resConfs.Data[1].Version)
}
