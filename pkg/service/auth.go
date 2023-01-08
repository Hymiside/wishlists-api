package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Hymiside/wishlists-api/pkg/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/Hymiside/wishlists-api/pkg/repository"
	"github.com/golang-jwt/jwt"
)

var (
	signingKey = []byte("qrkjk#4#%35FSFJlja#4353KSFjH")
	tokenTTL   = 12 * time.Hour
)

type Claims struct {
	jwt.StandardClaims
	UserId string
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user models.User) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthService) GenerateToken(email, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthService) ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return signingKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("error to parse jwt-token: %v", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
