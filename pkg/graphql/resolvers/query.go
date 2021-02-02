package graphql

import (
	"context"

	"quickeat/pkg/graphql/gqlgen"
	"quickeat/pkg/graphql/models"
)

type queryResolver struct{}

func NewQueryResolver() gqlgen.QueryResolver {
	return new(queryResolver)
}

func (q queryResolver) Dish(ctx context.Context, name string, category []string) ([]*models.Dish, error) {
	dish := models.NewDish()
	return dish, nil
}
