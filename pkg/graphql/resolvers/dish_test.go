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

func TestDishResolver_Category(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	categoryService := mock.NewMockCategoryService(ctrl)
	srvc := services.All{
		Category: categoryService,
	}
	dishResolver := NewDishResolver(srvc)

	t.Run("success", func(t *testing.T) {
		dish := models.Dish{
			ID: 200,
		}
		category := entity.Category{
			Id:   1,
			Name: "Brasileira",
		}
		categoryService.EXPECT().GetByDish(gomock.Any(), 200).Return(&category, nil)

		res, err := dishResolver.Category(ctx, &dish)
		require.Nil(t, err)
		require.Equal(t, 1, res.ID)
		require.Equal(t, "Brasileira", res.Name)
	})

	t.Run("failed", func(t *testing.T) {
		dish := models.Dish{
			ID: 500,
		}
		expectedErr := errors.New("failed")
		categoryService.EXPECT().GetByDish(gomock.Any(), 500).Return(nil, expectedErr)

		res, err := dishResolver.Category(ctx, &dish)
		require.Empty(t, res)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("empty", func(t *testing.T) {
		dish := models.Dish{
			ID: 400,
		}
		categoryService.EXPECT().GetByDish(gomock.Any(), 400).Return(nil, nil)

		res, err := dishResolver.Category(ctx, &dish)
		require.Empty(t, res)
		require.NoError(t, err)
	})
}
