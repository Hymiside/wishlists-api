package service

import (
	"errors"

	"github.com/Hymiside/wishlists-api/pkg/models"
	"github.com/Hymiside/wishlists-api/pkg/repository"
)

type Service struct {
	Authorization
}

var (
	ErrCreateJWT   = errors.New("error create jwt-token")
	ErrInvalidPwd  = errors.New("invalid password")
	ErrTokenClaims = errors.New("token claims are not of type *tokenClaims")
	ErrParseJWT    = errors.New("error to parse jwt-token")
	ErrSignMethod  = errors.New("invalid signing method")
	ErrHashPwd     = errors.New("error to hash password")
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (string, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
