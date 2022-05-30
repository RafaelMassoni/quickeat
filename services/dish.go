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
	UpdateDishName(ctx context.Context, dish *entity.Dish) error
	UpdateDishPrepTime(ctx context.Context, dish *entity.Dish) error
	UpdateDishPrice(ctx context.Context, dish *entity.Dish) error
	UpdateDishCategory(ctx context.Context, dish *entity.Dish) error
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

func (d dishService) UpdateDishName(ctx context.Context, dish *entity.Dish) error {
	query := `
		UPDATE pratos
		SET nome = :nome
		WHERE id = :id
	`

	_, err := d.db.NamedExecContext(ctx, query, dish)
	if err != nil {
		return err
	}

	return nil
}

func (d dishService) UpdateDishPrepTime(ctx context.Context, dish *entity.Dish) error {
	query := `
		UPDATE pratos
		SET tempo_de_preparo = :tempo_de_preparo
		WHERE id = :id
	`

	_, err := d.db.NamedExecContext(ctx, query, dish)
	if err != nil {
		return err
	}

	return nil
}

func (d dishService) UpdateDishPrice(ctx context.Context, dish *entity.Dish) error {
	query := `
		UPDATE pratos
		SET preco = :preco
		WHERE id = :id
		`

	_, err := d.db.NamedExecContext(ctx, query, dish)
	if err != nil {
		return err
	}

	return nil
}
func (d dishService) UpdateDishCategory(ctx context.Context, dish *entity.Dish) error {
	query := `
		UPDATE pratos
		SET id_categoria = :id_categoria
		WHERE id = :id
	`

	_, err := d.db.NamedExecContext(ctx, query, dish)
	if err != nil {
		return err
	}

	return nil
}

func (d dishService) CreateDish(ctx context.Context, dish *entity.Dish) error {

	query := `
		INSERT INTO pratos (id_categoria, nome, preco, tempo_de_preparo)
		VALUES (:id_categoria, :nome, :preco, :tempo_de_preparo)
	`

	result, err := d.db.NamedExecContext(ctx, query, dish)
	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	dish.Id = int(id)
	return nil
}
