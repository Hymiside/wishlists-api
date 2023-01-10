package service

import (
	"errors"

	"github.com/Hymiside/wishlists-api/pkg/models"
	"github.com/Hymiside/wishlists-api/pkg/repository"
)

var (
	ErrReadImage   = errors.New("error read image")
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

type Profile interface {
	GetProfile(userId string) (map[string]string, error)
}

type Service struct {
	Authorization
	Profile
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Profile:       NewProfile(repos.PersonalCabinet),
	}
}
