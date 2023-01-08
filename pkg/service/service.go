package service

import (
	"github.com/Hymiside/wishlists-api/pkg/models"
	"github.com/Hymiside/wishlists-api/pkg/repository"
)

type Service struct {
	Authorization
}

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
