package graphql

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"quickeat/pkg/entity"
	"quickeat/pkg/graphql/models"
	"quickeat/services"
	mock "quickeat/test"
)

func TestCategoryResolver_Dishes(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	dishService := mock.NewMockDishService(ctrl)
	srvc := services.All{
		Dish: dishService,
	}
	categoryResolver := NewCategoryResolver(srvc)

	t.Run("success", func(t *testing.T) {
		category := models.Category{
			ID: 1,
		}
		dishes := []*entity.Dish{
			{
				Id:         200,
				CategoryID: &[]int{1}[0],
				Name:       "Feijuca",
				Price:      20,
				CookTime:   50,
			},
			{
				Id:         201,
				CategoryID: &[]int{1}[0],
				Name:       "Macarrao",
				Price:      25,
				CookTime:   30,
			},
		}
		dishService.EXPECT().GetByCategory(gomock.Any(), 1).Return(dishes, nil)

		res, err := categoryResolver.Dishes(ctx, &category)
		require.Nil(t, err)
		require.Len(t, res, 2)
		require.Equal(t, "Feijuca", res[0].Name)
		require.Equal(t, "Macarrao", res[1].Name)
	})

	t.Run("failed", func(t *testing.T) {
		category := models.Category{
			ID: 500,
		}
		expectedErr := errors.New("failed")
		dishService.EXPECT().GetByCategory(gomock.Any(), 500).Return(nil, expectedErr)

		res, err := categoryResolver.Dishes(ctx, &category)
		require.Empty(t, res)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("empty", func(t *testing.T) {
		category := models.Category{
			ID: 400,
		}
		dishService.EXPECT().GetByCategory(gomock.Any(), 400).Return(nil, nil)

		res, err := categoryResolver.Dishes(ctx, &category)
		require.Empty(t, res)
		require.NoError(t, err)
	})
}
