package graphql

import (
	"context"

	"quickeat/pkg/graphql/gqlgen"
	"quickeat/pkg/graphql/models"
	"quickeat/services"
)

type queryResolver struct {
	services services.All
}

func NewQueryResolver(services services.All) gqlgen.QueryResolver {
	return queryResolver{services: services}
}

func (q queryResolver) Category(ctx context.Context, id int) (*models.Category, error) {
	category := models.NewCategory()
	return category, nil
}

func (q queryResolver) Dish(ctx context.Context, name string, category []string) ([]*models.Dish, error) {
	dish := models.NewDish()
	return dish, nil
}
