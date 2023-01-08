package repository

import (
	"database/sql"
	"github.com/Hymiside/wishlists-api/pkg/models"
	_ "github.com/lib/pq"
)

type Authorization interface {
	CreateUser(school models.User) (string, error)
	GetUser(email string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
