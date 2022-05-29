package graphql

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"quickeat/pkg/entity"
	"quickeat/services"
	mock "quickeat/test"
)

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
}
