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
	DeleteDishByName(ctx context.Context, DishName string) error
	DeleteDishById(ctx context.Context, DishId int) error
	UpdateDishName(ctx context.Context, DishId int, NewDishName string) error
	UpdateDishPrepTime(ctx context.Context, DishId int, NewDishPrepTime int) error
	UpdateDishPrice(ctx context.Context, DishId int, NewDishPrice int) error
	UpdateDishCategory(ctx context.Context, DishId int, NewDishCategory string) error
	CreateDish(ctx context.Context, dish *entity.Dish) error
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

func (d dishService) DeleteDishByName(ctx context.Context, DishName string) error {
	result := make([]*entity.Dish, 0)

	query := `
		DELETE FROM pratos WHERE nome = ?
	`

	err := d.db.Select(&result, query, DishName)
	if err != nil {
		return err
	}

	return nil
}

func (d dishService) DeleteDishById(ctx context.Context, DishId int) error {
	result := make([]*entity.Dish, 0)

	query := `
		DELETE FROM pratos WHERE id = ?
	`

	err := d.db.Select(&result, query, DishId)
	if err != nil {
		return err
	}

	return nil
}

func (d dishService) UpdateDishName(ctx context.Context, DishId int, NewDishName string) error {
	result := make([]*entity.Dish, 0)

	query := `
		UPDATE pratos
		SET nome = NewDishName
		WHERE DishId = ?;
	`

	err := d.db.Select(&result, query, DishId)
	if err != nil {
		return err
	}

	return nil
}

func (d dishService) UpdateDishPrepTime(ctx context.Context, DishId int, NewDishPrepTime int) error {
	result := make([]*entity.Dish, 0)

	query := `
		UPDATE pratos
		SET tempo_de_preparo = NewDishPrepTime
		WHERE DishId = ?;
	`

	err := d.db.Select(&result, query, DishId)
	if err != nil {
		return err
	}

	return nil
}

func (d dishService) UpdateDishPrice(ctx context.Context, DishId int, NewDishPrice int) error {
	result := make([]*entity.Dish, 0)

	query := `
		UPDATE pratos
		SET preco = NewDishPrice
		WHERE DishId = ?;	`

	err := d.db.Select(&result, query, DishId)
	if err != nil {
		return err
	}

	return nil
}
func (d dishService) UpdateDishCategory(ctx context.Context, DishId int, NewDishCategory string) error {
	result := make([]*entity.Dish, 0)

	query := `
		UPDATE pratos
		SET id_categoria = NewDishCategory
		WHERE DishId = ?;
	`

	err := d.db.Select(&result, query, DishId)
	if err != nil {
		return err
	}

	return nil
}

func (d dishService) CreateDish(ctx context.Context, dish *entity.Dish) error {
	result := make([]*entity.Dish, 0)

	query := `
		INSERT INTO pratos (id, id_categoria, nome, preco, tempo_de_preparo)
		VALUES (:id, :category, :name, :price, :tempPrep)
	`

	err := d.db.Select(&result, query, dish)
	if err != nil {
		return err
	}

	return nil
}
