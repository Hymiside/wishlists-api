package models

import "github.com/google/uuid"

type ConfigServer struct {
	Port string
	Host string
}

type ConfigRepository struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type User struct {
	Id          uuid.UUID `json:"user_id,omitempty" db:"id"`
	Name        string    `json:"name,omitempty" db:"name"`
	Nickname    string    `json:"nickname,omitempty" db:"nickname"`
	Email       string    `json:"email,omitempty" db:"email"`
	Password    string    `json:"password,omitempty" db:"password_hash"`
	Salt        string    `json:"-,omitempty" db:"salt"`
	PhoneNumber string    `json:"phone_number,omitempty" db:"phone_number"`
	ImageURL    string    `json:"image_url,omitempty" db:"image_url"`
}
