package sqlite

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.com/ulexxander/remoconf/storage"
)

type UsersStore struct {
	db *sqlx.DB
}

func NewUsersStore(db *sqlx.DB) *UsersStore {
	return &UsersStore{db}
}

const migrationUp = `
CREATE TABLE users (
	id integer PRIMARY KEY AUTOINCREMENT,
  login text NOT NULL,
  password text NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp
)
`

const migrationDown = `
DROP TABLE IF EXISTS users
`

func (us *UsersStore) Migrate() error {
	_, err := us.db.Exec(migrationDown)
	if err != nil {
		return fmt.Errorf("migration down: %w", err)
	}
	_, err = us.db.Exec(migrationUp)
	if err != nil {
		return fmt.Errorf("migration up: %w", err)
	}
	return nil
}

const getByIDQuery = `SELECT * FROM users WHERE id = $1`

func (us *UsersStore) GetByID(id int) (*storage.User, error) {
	var u storage.User
	if err := us.db.Get(&u, getByIDQuery, id); err != nil {
		return nil, err
	}
	return &u, nil
}

const createQuery = `INSERT INTO users (login, password)
VALUES ($1, $2)
RETURNING id`

func (us *UsersStore) Create(p storage.UserCreateParams) (int, error) {
	var id int
	if err := us.db.Get(&id, createQuery, p.Login, p.Password); err != nil {
		return 0, err
	}
	return id, nil
}
