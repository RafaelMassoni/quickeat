package services

import "github.com/jmoiron/sqlx"

type All struct {
	Dish     DishService
	Category CategoryService
}

func NewServices(db *sqlx.DB) All {
	return All{
		Dish:     NewDishService(db),
		Category: NewCategoryService(db),
	}
}
