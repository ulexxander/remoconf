package sqlite

import (
	"database/sql"
	"fmt"

	"gitlab.com/ulexxander/remoconf/storage"
)

type UsersStore struct {
	db *sql.DB
}

func NewUsersStore(db *sql.DB) *UsersStore {
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
	row := us.db.QueryRow(getByIDQuery, id)
	var u storage.User
	if err := row.Scan(
		&u.ID,
		&u.Login,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &u, nil
}

const createQuery = `INSERT INTO users (login, password)
VALUES ($1, $2)
RETURNING id`

func (us *UsersStore) Create(p storage.UserCreateParams) (int, error) {
	row := us.db.QueryRow(createQuery, p.Login, p.Password)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
