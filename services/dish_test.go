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

func TestDishService_DeleteDishByName(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("DELETE FROM pratos WHERE nome = ?")

	t.Run("failed", func(t *testing.T) {
		dishName := "test_sucess"
		expectedErr := errors.New("Name not found")

		mock.ExpectQuery(query).
			WithArgs(dishName).
			WillReturnError(expectedErr)

		err := srvc.DeleteDishByName(ctx, dishName)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		dishName := "test_failed"

		mock.ExpectQuery(query).
			WithArgs(dishName).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_categoria", "nome", "preco", "tempo_de_preparo"}).
				AddRow(1, 1, "dishName", 12, 20))

		err := srvc.DeleteDishByName(ctx, dishName)
		require.NoError(t, err)
	})
}

func TestDishService_DeleteDishById(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("DELETE FROM pratos WHERE id = ?")

	t.Run("failed", func(t *testing.T) {
		dishId := 1
		expectedErr := errors.New("Id not found")

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		err := srvc.DeleteDishById(ctx, dishId)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		dishId := 1

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_categoria", "nome", "preco", "tempo_de_preparo"}).
				AddRow(1, 1, "dishName", 12, 20))

		err := srvc.DeleteDishById(ctx, dishId)
		require.NoError(t, err)
	})
}

func TestDishService_UpdateDishName(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("UPDATE pratos SET nome = NewDishName WHERE DishId = ?;")

	t.Run("failed", func(t *testing.T) {
		dishId := 1
		dishName := "test"
		expectedErr := errors.New("Id not found")

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		err := srvc.UpdateDishName(ctx, dishId, dishName)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		dishId := 1
		dishName := "test"

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_categoria", "nome", "preco", "tempo_de_preparo"}).
				AddRow(dishId, 1, dishName, 12, 20))

		err := srvc.UpdateDishName(ctx, dishId, dishName)
		require.NoError(t, err)
	})
}

func TestDishService_UpdateDishCategory(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("UPDATE pratos SET id_categoria = NewDishCategory WHERE DishId = ?;")

	t.Run("failed", func(t *testing.T) {
		dishId := 1
		dishCategory := "test"
		expectedErr := errors.New("Id not found")

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		err := srvc.UpdateDishCategory(ctx, dishId, dishCategory)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		dishId := 1
		dishCategory := "test"

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_categoria", "nome", "preco", "tempo_de_preparo"}).
				AddRow(dishId, 1, dishCategory, 12, 20))

		err := srvc.UpdateDishCategory(ctx, dishId, dishCategory)
		require.NoError(t, err)
	})
}

func TestDishService_UpdateDishPrice(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("UPDATE pratos SET preco = NewDishPrice WHERE DishId = ?;")

	t.Run("failed", func(t *testing.T) {
		dishId := 1
		dishPrice := 10
		expectedErr := errors.New("Id not found")

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		err := srvc.UpdateDishPrice(ctx, dishId, dishPrice)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		dishId := 1
		dishPrice := 2

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_categoria", "nome", "preco", "tempo_de_preparo"}).
				AddRow(dishId, 1, dishPrice, 12, 20))

		err := srvc.UpdateDishPrice(ctx, dishId, dishPrice)
		require.NoError(t, err)
	})
}

func TestDishService_UpdateDishPrepTime(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("UPDATE pratos SET tempo_de_preparo = NewDishPrepTime WHERE DishId = ?;")

	t.Run("failed", func(t *testing.T) {
		dishId := 1
		dishPrepTime := 2
		expectedErr := errors.New("Id not found")

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		err := srvc.UpdateDishPrepTime(ctx, dishId, dishPrepTime)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		dishId := 1
		dishPrepTime := 2

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_categoria", "nome", "preco", "tempo_de_preparo"}).
				AddRow(dishId, 1, dishPrepTime, 12, 20))

		err := srvc.UpdateDishPrepTime(ctx, dishId, dishPrepTime)
		require.NoError(t, err)
	})
}

func TestDishService_CreateDish(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("INSERT INTO pratos (id, id_categoria, nome, preco, tempo_de_preparo) VALUES (:id, :category, :name, :price, :tempPrep)")

	t.Run("failed", func(t *testing.T) {
		dishId := 1
		dishCategory := 1
		dishName := "nameTest"
		dishprice := 10
		dishPrepTime := 2
		expectedErr := errors.New("Id not found")

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		err := srvc.CreateDish(ctx, dishId, dishCategory, dishName, dishprice, dishPrepTime)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		dishId := 1
		dishCategory := 1
		dishName := "nameTest"
		dishprice := 10
		dishPrepTime := 2

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_categoria", "nome", "preco", "tempo_de_preparo"}).
				AddRow(dishId, dishCategory, dishName, dishprice, dishPrepTime))

		err := srvc.CreateDish(ctx, dishId, dishCategory, dishName, dishprice, dishPrepTime)
		require.NoError(t, err)
	})
}
