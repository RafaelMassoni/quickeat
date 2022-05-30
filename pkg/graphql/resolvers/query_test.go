package graphql

import (
	"context"
	"errors"
	"testing"

	"quickeat/pkg/entity"
	"quickeat/services"
	mock "quickeat/test"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestQueryResolver_Category(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	categoryService := mock.NewMockCategoryService(ctrl)
	srvc := services.All{
		Category: categoryService,
	}
	query := NewQueryResolver(srvc)

	t.Run("success with id", func(t *testing.T) {
		Id := 200
		category := []*entity.Category{
			{
				Id:   200,
				Name: "brasileira",
			},
		}

		categoryService.EXPECT().Get(gomock.Any(), gomock.Eq(&Id)).Return(category, nil)
		res, err := query.Category(ctx, &Id)

		require.Nil(t, err)
		require.Equal(t, "brasileira", res[0].Name)
		require.Equal(t, 200, res[0].ID)
	})

	t.Run("success get all", func(t *testing.T) {
		categories := []*entity.Category{
			{
				Id:   200,
				Name: "brasileira",
			},
			{
				Id:   201,
				Name: "japonesa",
			},
			{
				Id:   202,
				Name: "indiana",
			},
		}

		categoryService.EXPECT().Get(gomock.Any(), nil).Return(categories, nil)
		res, err := query.Category(ctx, nil)

		require.Nil(t, err)
		require.Len(t, res, 3)
		require.Equal(t, "brasileira", res[0].Name)
		require.Equal(t, 200, res[0].ID)
		require.Equal(t, "indiana", res[2].Name)
		require.Equal(t, 202, res[2].ID)
	})

	t.Run("empty", func(t *testing.T) {
		categoryService.EXPECT().Get(gomock.Any(), nil).Return(nil, nil)
		res, err := query.Category(ctx, nil)

		require.Nil(t, err)
		require.Empty(t, res)
	})

	t.Run("failed with id", func(t *testing.T) {
		id := 500
		expectedErr := errors.New("service failed")
		categoryService.EXPECT().Get(gomock.Any(), &id).Return(nil, expectedErr)
		res, err := query.Category(ctx, &id)

		require.Empty(t, res)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("failed without id", func(t *testing.T) {
		expectedErr := errors.New("service failed")
		categoryService.EXPECT().Get(gomock.Any(), nil).Return(nil, expectedErr)
		res, err := query.Category(ctx, nil)

		require.Empty(t, res)
		require.EqualError(t, err, expectedErr.Error())
	})
}

func TestQueryResolver_Dish(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	dishService := mock.NewMockDishService(ctrl)
	srvc := services.All{
		Dish: dishService,
	}
	query := NewQueryResolver(srvc)

	t.Run("success with id", func(t *testing.T) {
		Id := 200
		dish := entity.Dish{
			Id:       200,
			Name:     "macarronada",
			Price:    15,
			CookTime: 25,
		}

		dishService.EXPECT().Get(gomock.Any(), Id).Return(&dish, nil)
		res, err := query.Dish(ctx, &Id)

		require.Nil(t, err)
		require.Equal(t, "macarronada", res[0].Name)
		require.Equal(t, 200, res[0].ID)
	})

	t.Run("empty with id", func(t *testing.T) {
		Id := 400

		dishService.EXPECT().Get(gomock.Any(), Id).Return(nil, nil)
		res, err := query.Dish(ctx, &Id)

		require.Nil(t, err)
		require.Empty(t, res)
	})

	t.Run("success get all", func(t *testing.T) {
		dishes := []*entity.Dish{
			{
				Id:       200,
				Name:     "macarronada",
				Price:    15,
				CookTime: 25,
			},
			{
				Id:       201,
				Name:     "feijoada",
				Price:    20,
				CookTime: 45,
			},
			{
				Id:       202,
				Name:     "bife de figado",
				Price:    15,
				CookTime: 30,
			},
		}

		dishService.EXPECT().GetAll(gomock.Any()).Return(dishes, nil)
		res, err := query.Dish(ctx, nil)

		require.Nil(t, err)
		require.Len(t, res, 3)
		require.Equal(t, "macarronada", res[0].Name)
		require.Equal(t, "bife de figado", res[2].Name)
		require.Equal(t, 201, res[1].ID)
	})

	t.Run("empty get all", func(t *testing.T) {
		dishService.EXPECT().GetAll(gomock.Any()).Return(nil, nil)
		res, err := query.Dish(ctx, nil)

		require.Nil(t, err)
		require.Empty(t, res)
	})

	t.Run("failed with id", func(t *testing.T) {
		id := 500
		expectedErr := errors.New("service failed")
		dishService.EXPECT().Get(gomock.Any(), id).Return(nil, expectedErr)
		res, err := query.Dish(ctx, &id)

		require.Empty(t, res)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("failed without id", func(t *testing.T) {
		expectedErr := errors.New("service failed")
		dishService.EXPECT().GetAll(gomock.Any()).Return(nil, expectedErr)
		res, err := query.Dish(ctx, nil)

		require.Empty(t, res)
		require.EqualError(t, err, expectedErr.Error())
	})
}
