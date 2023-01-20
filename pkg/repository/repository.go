package repository

import (
	"database/sql"
	"errors"

	"github.com/Hymiside/wishlists-api/pkg/models"
	_ "github.com/lib/pq"
)

var (
	ErrUserNotFound       = errors.New("user not found")          // Не нашел пользователя в БД
	ErrItemsNotFound      = errors.New("items not found")         // Объект не наден в БД
	ErrScanItems          = errors.New("error scan items")        // Ошибка сканирования объекта row после запроса в БД
	ErrQueryItems         = errors.New("error get items from db") // Ошибка при запросе в БД
	ErrUniqueKeyViolation = errors.New("key unique violation")    // Нарушение уникальности в таблице
	ErrCreateUser         = errors.New("error create user")       // Не удалось создать пользователя, по никому неизвестной причине
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	GetUser(email string) (models.User, error)
}

type PersonalCabinet interface {
	GetProfile(userId string) (map[string]string, error)
	GetWishes(userId string) ([]models.Wish, error)
}

type Repository struct {
	Authorization
	PersonalCabinet
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization:   NewAuthPostgres(db),
		PersonalCabinet: NewProfilePostgres(db),
	}
}
