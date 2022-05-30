package graphql

import (
	"context"
	"errors"

	"quickeat/pkg/entity"
	"quickeat/pkg/graphql/gqlgen"
	"quickeat/pkg/graphql/models"
	"quickeat/services"
)

type mutationResolver struct {
	services services.All
}

func NewMutationResolver(services services.All) gqlgen.MutationResolver {
	return mutationResolver{services: services}
}

func (m mutationResolver) CreateDish(ctx context.Context, input models.CreateDishInput) (bool, error) {
	if input.Name == "" {
		return false, errors.New("invalid name")
	}

	if input.Price == 0 {
		return false, errors.New("invalid price")
	}

	if input.CookTime == 0 {
		return false, errors.New("invalid cook time")
	}

	if input.Category == 0 {
		return false, errors.New("invalid category")
	}

	dish := entity.Dish{
		CategoryID: &input.Category,
		Name:       input.Name,
		Price:      input.Price,
		CookTime:   input.CookTime,
	}

	err := m.services.Dish.CreateDish(ctx, &dish)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m mutationResolver) CreateCategory(ctx context.Context, name string) (bool, error) {
	return true, nil
}
