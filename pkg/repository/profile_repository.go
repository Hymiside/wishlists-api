package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Hymiside/wishlists-api/pkg/models"
	"time"
)

type ProfilePostgres struct {
	db *sql.DB
}

func NewProfilePostgres(db *sql.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

func (p *ProfilePostgres) GetProfile(userId string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	request := fmt.Sprintf("select id, name, nickname, image_url from users where id = $1")

	row := p.db.QueryRowContext(ctx, request, userId)
	if row.Err() != nil {
		return nil, ErrQueryItems
	}

	var (
		id, name, nickname string
		imageUrl           sql.NullString
	)
	if err := row.Scan(&id, &name, &nickname, &imageUrl); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, ErrScanItems
	}

	// В таблице #subscribers_users, столбец #user_id считает подписчиков,
	// а столбец #user_id_sub считает подписки

	requestSub1 := fmt.Sprintf("select count(id) from subscribes_users where user_id = $1")
	requestSub2 := fmt.Sprintf("select count(id) from subscribes_users where user_id_sub = $1")

	var numSubscribers, numSubscriptions string
	if err := p.db.QueryRowContext(ctx, requestSub1, userId).Scan(&numSubscribers); err != nil {
		return nil, ErrScanItems
	}
	if err := p.db.QueryRowContext(ctx, requestSub2, userId).Scan(&numSubscriptions); err != nil {
		return nil, ErrScanItems
	}

	res := map[string]string{
		"user_id":       id,
		"name":          name,
		"nickname":      nickname,
		"subscribes":    numSubscribers,
		"subscriptions": numSubscriptions,
		"image_base64":  imageUrl.String,
	}

	return res, nil
}

func (p *ProfilePostgres) GetWishes(userId string) ([]models.Wish, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	request := fmt.Sprintf("select * from wishes where user_id = $1")

	rows, err := p.db.QueryContext(ctx, request, userId)
	if err != nil {
		return nil, ErrQueryItems
	}

	var wishes []models.Wish
	for rows.Next() {
		var wish models.Wish

		if err = rows.Scan(&wish.Id, &wish.UserId, &wish.Title, &wish.Description, &wish.Price,
			&wish.Link, &wish.ImageURL); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, ErrItemsNotFound
			}
			return nil, ErrScanItems
		}
		wishes = append(wishes, wish)
	}
	return wishes, nil
}

func (p *ProfilePostgres) CreateWish(wish models.Wish) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	request := fmt.Sprintf("insert into wishes (id, user_id, title, description, price, link, image_url) values ($1, $2, $3, $4, $5, $6, $7)" +
		"returning id;")

	var wishId string
	if err := p.db.QueryRowContext(ctx, request, wish.Id, wish.UserId, wish.Title, wish.Description,
		wish.Price, wish.Link, wish.ImageURL).Scan(&wishId); err != nil {
		return "", ErrCreateItem
	}
	return wishId, nil
}
