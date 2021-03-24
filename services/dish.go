//go:generate go run -mod=mod github.com/golang/mock/mockgen -package=mock -source=$GOFILE -destination=../test/dishservice.go
package services

import (
	"context"

	"github.com/jmoiron/sqlx"

	"quickeat/pkg/entity"
)

type DishService interface {
	Get(ctx context.Context, id int) (*entity.Dish, error)
	GetAll(ctx context.Context) ([]*entity.Dish, error)
	GetByCategory(ctx context.Context, categoryID int) ([]*entity.Dish, error)
}

type dishService struct {
	db *sqlx.DB
}

func NewDishService(db *sqlx.DB) DishService {
	return dishService{db: db}
}

func (d dishService) Get(ctx context.Context, id int) (*entity.Dish, error) {
	var result []*entity.Dish
	query := `
		SELECT * FROM pratos WHERE id = ?
	`

	err := d.db.Select(&result, query, id)
	if err != nil {
		return nil, err
	}

	return result[0], nil
}

func (d dishService) GetAll(ctx context.Context) ([]*entity.Dish, error) {
	result := make([]*entity.Dish, 0)

	query := `
		SELECT * FROM pratos 
	`

	err := d.db.Select(&result, query)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (d dishService) GetByCategory(ctx context.Context, categoryID int) ([]*entity.Dish, error) {
	result := make([]*entity.Dish, 0)

	query := `
		SELECT * FROM pratos WHERE id_categoria = ?
	`

	err := d.db.Select(&result, query, categoryID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
