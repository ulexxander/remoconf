package sqlite_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/ulexxander/remoconf/storage"
	"gitlab.com/ulexxander/remoconf/storage/sqlite"
	"gitlab.com/ulexxander/remoconf/testutil"
)

func TestConfigsStore(t *testing.T) {
	db := testutil.SetupDBTest(t)
	us := sqlite.NewUsersStore(db)
	ps := sqlite.NewProjectsStore(db)
	cs := sqlite.NewConfigsStore(db)

	t.Run("no configs", func(t *testing.T) {
		p, err := cs.GetAll()
		require.NoError(t, err)
		require.Nil(t, p)
	})

	version := 1
	content := "my cfg"
	user, _ := us.Create(storage.UserCreateParams{})
	project, _ := ps.Create(storage.ProjectCreateParams{CreatedBy: user.ID})

	var configID int

	t.Run("creating config", func(t *testing.T) {
		t.Run("fails project does not exist", func(t *testing.T) {
			_, err := cs.Create(storage.ConfigCreateParams{
				ProjectID: 4151,
				Version:   version,
				Content:   content,
				CreatedBy: user.ID,
			})
			require.Error(t, err)
		})

		t.Run("fails user does not exist", func(t *testing.T) {
			_, err := cs.Create(storage.ConfigCreateParams{
				ProjectID: project.ID,
				Version:   version,
				Content:   content,
				CreatedBy: 23455,
			})
			require.Error(t, err)
		})

		t.Run("succeeds", func(t *testing.T) {
			created, err := cs.Create(storage.ConfigCreateParams{
				ProjectID: project.ID,
				Version:   version,
				Content:   content,
				CreatedBy: user.ID,
			})
			require.NoError(t, err)
			configID = created.ID
		})
	})

	t.Run("created config listed", func(t *testing.T) {
		c, err := cs.GetAll()
		require.NoError(t, err)
		require.Len(t, c, 1)

		require.Equal(t, configID, c[0].ID)
		require.Equal(t, project.ID, c[0].ProjectID)
		require.Equal(t, version, c[0].Version)
		require.Equal(t, content, c[0].Content)
		require.NotZero(t, c[0].CreatedAt)
		require.Equal(t, user.ID, c[0].CreatedBy)
	})
}
