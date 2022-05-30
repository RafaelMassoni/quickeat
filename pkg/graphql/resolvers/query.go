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

func (q queryResolver) Category(ctx context.Context, id *int) ([]*models.Category, error) {
	c, err := q.services.Category.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return models.NewCategory(c...), nil
}

func (q queryResolver) Dish(ctx context.Context, id *int) ([]*models.Dish, error) {
	if id != nil {
		d, err := q.services.Dish.Get(ctx, *id)
		if err != nil {
			return nil, err
		}
		return models.NewDish(d), nil
	}

	d, err := q.services.Dish.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return models.NewDish(d...), nil
}
