package storage

type UsersStore interface {
	GetByID(id int) (*User, error)
	Create(p UserCreateParams) (*CreatedItem, error)
}

type ProjectsStore interface {
	GetAll() ([]Project, error)
	Create(p ProjectCreateParams) (*CreatedItem, error)
}

type ConfigsStore interface {
	GetAll() ([]Config, error)
	Create(p ConfigCreateParams) (*CreatedItem, error)
}
