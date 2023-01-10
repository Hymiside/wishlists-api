package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

	var id, name, nickname, imageUrl string
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
		"image_base64":  imageUrl,
	}

	return res, nil
}
