package services

import (
	"context"

	"github.com/jmoiron/sqlx"

	"quickeat/pkg/entity"
)

type UserService interface {
	Get(ctx context.Context, id int) (*entity.User, error)
}

type userService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) UserService {
	return userService{db: db}
}

func (u userService) Get(ctx context.Context, id int) (*entity.User, error) {
	result := entity.User{}
	if err := u.db.Get(&result, "SELECT * FROM usuarios WHERE id = ?", id); err != nil {
		return nil, err
	}

	return &result, nil
}
