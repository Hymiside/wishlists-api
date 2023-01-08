package service

import (
	"time"

	"github.com/google/uuid"

	"github.com/Hymiside/wishlists-api/pkg/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/Hymiside/wishlists-api/pkg/repository"
	"github.com/golang-jwt/jwt"
)

var (
	signingKey = []byte("qrkjk#4#%35FSFJlja#4353KSFjH")
	tokenTTL   = 1460 * time.Hour
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
	var userId string

	user.Id = uuid.New().String()
	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		return "", ErrHashPwd
	}
	user.Password = passwordHash

	userId, err = a.repo.CreateUser(user)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (a *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := a.repo.GetUser(email)
	if err != nil {
		return "", err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", ErrInvalidPwd
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	var tokenString string
	tokenString, err = token.SignedString(signingKey)
	if err != nil {
		return "", ErrCreateJWT
	}
	return tokenString, nil
}

func (a *AuthService) ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrSignMethod
		}
		return signingKey, nil
	})
	if err != nil {
		return "", ErrParseJWT
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", ErrTokenClaims
	}
	return claims.UserId, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
