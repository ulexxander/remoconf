package sqlite_test

import (
	"testing"

	"gitlab.com/ulexxander/remoconf/storage/sqlite"
)

func TestUsersStore(t *testing.T) {
	us := sqlite.UsersStore{}

	t.Run("user does not exist", func(t *testing.T) {
		u, err := us.GetByID(1)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if u != nil {
			t.Fatalf("expected user to be nil, got %v", u)
		}
	})
}
