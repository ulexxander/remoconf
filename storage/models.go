package storage

import "time"

// models

type User struct {
	ID        int        `db:"id"`
	Login     string     `db:"login"`
	Password  string     `db:"password"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type Project struct {
	ID          int        `db:"id"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
	CreatedAt   time.Time  `db:"created_at"`
	CreatedBy   int        `db:"created_by"`
	UpdatedAt   *time.Time `db:"updated_at"`
	UpdatedBy   *int       `db:"updated_by"`
}

type Config struct {
	ID        int       `db:"id"`
	ProjectID uint      `db:"project_id"`
	Version   int       `db:"version"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	CreatedBy int       `db:"created_by"`
}

// params

type UserCreateParams struct {
	Login    string
	Password string
}
