package sqlite_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/ulexxander/remoconf/storage"
	"gitlab.com/ulexxander/remoconf/storage/sqlite"
	"gitlab.com/ulexxander/remoconf/testutil"
)

func TestProjectsStore(t *testing.T) {
	db := testutil.SetupDBTest(t)
	ps := sqlite.NewProjectsStore(db)
	us := sqlite.NewUsersStore(db)

	t.Run("no projects", func(t *testing.T) {
		p, err := ps.GetAll()
		require.NoError(t, err)
		require.Nil(t, p)
	})

	title := "my proj"
	description := "some desc"
	createdBy, _ := us.Create(storage.UserCreateParams{})

	var projectID int

	t.Run("creating project", func(t *testing.T) {
		t.Run("fails user does not exist", func(t *testing.T) {
			_, err := ps.Create(storage.ProjectCreateParams{
				Title:       title,
				Description: description,
				CreatedBy:   1452,
			})
			require.Error(t, err)
		})

		t.Run("succeeds", func(t *testing.T) {
			id, err := ps.Create(storage.ProjectCreateParams{
				Title:       title,
				Description: description,
				CreatedBy:   createdBy,
			})
			require.NoError(t, err)
			projectID = id
		})
	})

	t.Run("created project listed", func(t *testing.T) {
		p, err := ps.GetAll()
		require.NoError(t, err)
		require.Len(t, p, 1)

		require.Equal(t, projectID, p[0].ID)
		require.Equal(t, title, p[0].Title)
		require.Equal(t, description, p[0].Description)
		require.NotZero(t, p[0].CreatedAt)
		require.Equal(t, createdBy, p[0].CreatedBy)
		require.Nil(t, p[0].UpdatedAt)
		require.Nil(t, p[0].UpdatedBy)
	})
}
