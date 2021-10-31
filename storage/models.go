package storage

import "time"

// models

type User struct {
	ID        int
	Login     string
	Password  string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type Project struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	CreatedBy   int
	UpdatedAt   *time.Time
	UpdatedBy   *int
}

type Config struct {
	ID        int
	ProjectID uint
	Version   int
	Content   string
	CreatedAt time.Time
	CreatedBy int
}

// params

type UserCreateParams struct {
	Login    string
	Password string
}
