package graphql

import (
	"context"

	"quickeat/pkg/graphql/gqlgen"
	"quickeat/pkg/graphql/models"
	"quickeat/services"
)

type dishResolver struct {
	services services.All
}

func NewDishResolver(s services.All) gqlgen.DishResolver {
	return dishResolver{services: s}
}

func (d dishResolver) Category(ctx context.Context, dish *models.Dish) (*models.Category, error) {
	categories, err := d.services.Category.GetByDish(ctx, dish.ID)
	if err != nil {
		return nil, err
	}
	if categories.IsEmpty() {
		return nil, nil
	}
	return models.NewCategory(categories)[0], nil
}
