package users

import (
	"github.com/pkg/errors"
	"gitlab.com/ulexxander/remoconf/storage"
	"golang.org/x/crypto/bcrypt"
)

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
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "hashing password")
	}
	created, err := s.users.Create(storage.UserCreateParams{
		Login:    p.Login,
		Password: string(passwordHashed),
	})
	if err != nil {
		return nil, errors.Wrap(err, "creating user in db")
	}
	return created, nil
}
