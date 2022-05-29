package services

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestDishService_Get(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("SELECT * FROM pratos WHERE id = ?")

	t.Run("success", func(t *testing.T) {
		dishId := 200

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_categoria", "nome", "preco", "tempo_de_preparo"}).
				AddRow(dishId, 1, "macarronada", 12, 20))

		res, err := srvc.Get(ctx, dishId)
		require.NoError(t, err)
		require.Equal(t, dishId, res.Id)
		require.Equal(t, "macarronada", res.Name)
	})

	t.Run("failed", func(t *testing.T) {
		dishId := 500
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		res, err := srvc.Get(ctx, dishId)
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestDishService_GetAll(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("SELECT * FROM pratos")

	t.Run("success", func(t *testing.T) {
		dishId1 := 200
		dishId2 := 220
		mock.ExpectQuery(query).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_categoria", "nome", "preco", "tempo_de_preparo"}).
				AddRow(dishId1, 1, "macarronada", 12, 20).
				AddRow(dishId2, 1, "carbonara", 30, 30))

		res, err := srvc.GetAll(ctx)
		require.NoError(t, err)
		require.Equal(t, dishId1, res[0].Id)
		require.Equal(t, "macarronada", res[0].Name)
		require.Equal(t, dishId2, res[1].Id)
		require.Equal(t, "carbonara", res[1].Name)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WillReturnError(expectedErr)

		res, err := srvc.GetAll(ctx)
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestDishService_GetByCategory(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("SELECT * FROM pratos WHERE id_categoria = ?")

	t.Run("success", func(t *testing.T) {
		categoryId := 1
		mock.ExpectQuery(query).
			WithArgs(categoryId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_categoria", "nome", "preco", "tempo_de_preparo"}).
				AddRow(200, categoryId, "macarronada", 12, 20))

		res, err := srvc.GetByCategory(ctx, categoryId)
		require.NoError(t, err)
		require.Equal(t, categoryId, *res[0].CategoryID)
		require.Equal(t, "macarronada", res[0].Name)
	})

	t.Run("failed", func(t *testing.T) {
		categoryId := 1
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WithArgs(categoryId).
			WillReturnError(expectedErr)

		res, err := srvc.GetByCategory(ctx, categoryId)
		require.Error(t, err)
		require.Nil(t, res)
	})
}
