package repository

import (
	"database/sql"
	"github.com/Hymiside/wishlists-api/pkg/models"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user models.User) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthPostgres) GetUser(email string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}
