package repository

import (
	"database/sql"
	"errors"

	"github.com/Hymiside/wishlists-api/pkg/models"
	_ "github.com/lib/pq"
)

var (
	ErrItemsNotFound      = errors.New("items not found")
	ErrScanItems          = errors.New("error scan items")
	ErrQueryItems         = errors.New("error get items from db")
	ErrUniqueKeyViolation = errors.New("key unique violation")
	ErrCreateUser         = errors.New("error create user")
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
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
