package sqlite

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/ulexxander/remoconf/storage"
)

type UsersStore struct {
	db *sqlx.DB
}

func NewUsersStore(db *sqlx.DB) *UsersStore {
	return &UsersStore{db}
}

const userGetByIDQuery = `SELECT * FROM users WHERE id = $1`

func (s *UsersStore) GetByID(id int) (*storage.User, error) {
	var item storage.User
	if err := s.db.Get(&item, userGetByIDQuery, id); err != nil {
		return nil, err
	}
	return &item, nil
}

const userCreateQuery = `INSERT INTO users (login, password)
VALUES ($1, $2)
RETURNING id`

func (s *UsersStore) Create(p storage.UserCreateParams) (*storage.CreatedItem, error) {
	var created storage.CreatedItem
	if err := s.db.Get(&created, userCreateQuery, p.Login, p.Password); err != nil {
		return &created, err
	}
	return &created, nil
}
