package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Hymiside/wishlists-api/pkg/models"
	"github.com/lib/pq"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	request := fmt.Sprintf("insert into users (id, name, nickname, email, password_hash) values ($1, $2, $3, $4, $5) returning id")

	var userId string
	if err := r.db.QueryRowContext(ctx, request, user.Id, user.Name, user.Nickname, user.Email, user.Password).Scan(&userId); err != nil {
		if err.(*pq.Error).Code == "23505" {
			return "", ErrUniqueKeyViolation
		}
		return "", ErrCreateUser
	}
	return userId, nil
}

func (r *AuthPostgres) GetUser(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	request := fmt.Sprintf("select id, password_hash from users where email = $1")

	var userId, passwordHash string
	row := r.db.QueryRowContext(ctx, request, email)
	if row.Err() != nil {
		return models.User{}, ErrQueryItems
	}

	if err := row.Scan(&userId, &passwordHash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, ErrItemsNotFound
		}
		return models.User{}, ErrScanItems
	}

	var user models.User
	user.Id, user.Password = userId, passwordHash

	return user, nil
}
