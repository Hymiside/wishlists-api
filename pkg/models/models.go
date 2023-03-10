package models

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
	Id          string `json:"user_id,omitempty" db:"id"`
	Name        string `json:"name,omitempty" db:"name"`
	Nickname    string `json:"nickname,omitempty" db:"nickname"`
	Email       string `json:"email,omitempty" db:"email"`
	Password    string `json:"password,omitempty" db:"password_hash"`
	PhoneNumber string `json:"phone_number,omitempty" db:"phone_number"`
	ImageURL    string `json:"image,omitempty" db:"image_url"`
}

type Wish struct {
	Id          string `json:"id,omitempty" db:"id"`
	UserId      string `json:"user_id,omitempty" db:"user_id"`
	Title       string `json:"title,omitempty" db:"title"`
	Description string `json:"description,omitempty" db:"description"`
	Price       int    `json:"price,omitempty" db:"price"`
	Link        string `json:"link,omitempty" db:"link"`
	ImageURL    string `json:"image_base64,omitempty" db:"image_url"`
}
