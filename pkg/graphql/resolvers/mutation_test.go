package graphql

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"quickeat/pkg/graphql/models"
	"quickeat/services"
	mock "quickeat/test"
)

func TestMutationResolver_CreateDish(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	dishService := mock.NewMockDishService(ctrl)
	srvc := services.All{
		Dish: dishService,
	}
	mutation := NewMutationResolver(srvc)

	t.Run("success", func(t *testing.T) {
		input := models.CreateDishInput{
			Name:     "Franguinho",
			Category: 1,
			Price:    15,
			CookTime: 20,
		}

		dishService.EXPECT().CreateDish(gomock.Any(), gomock.Any()).Return(nil)

		ok, err := mutation.CreateDish(ctx, input)
		require.True(t, ok)
		require.Nil(t, err)
	})

	t.Run("invalid name", func(t *testing.T) {
		input := models.CreateDishInput{}

		ok, err := mutation.CreateDish(ctx, input)
		require.False(t, ok)
		require.Error(t, err)
	})

	t.Run("invalid price", func(t *testing.T) {
		input := models.CreateDishInput{
			Name: "Franguinho",
		}

		ok, err := mutation.CreateDish(ctx, input)
		require.False(t, ok)
		require.Error(t, err)
	})

	t.Run("invalid cook time", func(t *testing.T) {
		input := models.CreateDishInput{
			Name:  "Franguinho",
			Price: 15,
		}

		ok, err := mutation.CreateDish(ctx, input)
		require.False(t, ok)
		require.Error(t, err)
	})

	t.Run("invalid category", func(t *testing.T) {
		input := models.CreateDishInput{
			Name:     "Franguinho",
			Price:    15,
			CookTime: 10,
		}

		ok, err := mutation.CreateDish(ctx, input)
		require.False(t, ok)
		require.Error(t, err)
	})
}
