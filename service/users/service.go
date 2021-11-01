package users

import "gitlab.com/ulexxander/remoconf/storage"

type Service struct {
	users storage.UsersStore
}

func NewService(us storage.UsersStore) *Service {
	return &Service{us}
}

func (s *Service) GetByID(id int) (*storage.User, error) {
	return s.users.GetByID(id)
}

func (s *Service) Create(p storage.UserCreateParams) (*storage.CreatedItem, error) {
	return s.users.Create(p)
}
