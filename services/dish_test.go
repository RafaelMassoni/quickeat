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
