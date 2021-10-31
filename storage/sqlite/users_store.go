package sqlite

import "gitlab.com/ulexxander/remoconf/storage"

type UsersStore struct {
}

func (us *UsersStore) GetByID(id int) (*storage.User, error) {
	return nil, nil
}

func (us *UsersStore) Create(p storage.UserCreateParams) (*storage.User, error) {
	return nil, nil
}
