package services

import (
	"context"
	"errors"
	"quickeat/pkg/entity"
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

	t.Run("success", func(t *testing.T) {
		dishName := "test_failed"

		mock.ExpectQuery(query).
			WithArgs(dishName).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_categoria", "nome", "preco", "tempo_de_preparo"}).
				AddRow(1, 1, "dishName", 12, 20))

		res := srvc.DeleteDishByName(ctx, dishName)
		require.Nil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		dishName := "test_sucess"
		expectedErr := errors.New("Name not found")

		mock.ExpectQuery(query).
			WithArgs(dishName).
			WillReturnError(expectedErr)

		err := srvc.DeleteDishByName(ctx, dishName)
		require.Error(t, err)
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

	t.Run("success", func(t *testing.T) {
		dishId := 1

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_categoria", "nome", "preco", "tempo_de_preparo"}).
				AddRow(1, 1, "dishName", 12, 20))

		res := srvc.DeleteDishById(ctx, dishId)
		require.Nil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		dishId := 1
		expectedErr := errors.New("Id not found")

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		err := srvc.DeleteDishById(ctx, dishId)
		require.Error(t, err)
	})
}

func TestDishService_UpdateDishName(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("UPDATE pratos SET nome = ? WHERE id = ?")

	dish := new(entity.Dish)
	dish.Id = 1
	dish.Name = "test"

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(dish.Name, dish.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))

		res := srvc.UpdateDishName(ctx, dish)
		require.Nil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("Id not found")

		mock.ExpectQuery(query).
			WithArgs(dish.Name, dish.Id).
			WillReturnError(expectedErr)

		err := srvc.UpdateDishName(ctx, dish)
		require.Error(t, err)
	})
}

func TestDishService_UpdateDishCategory(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("UPDATE pratos SET id_categoria = ? WHERE id = ?")

	dish := new(entity.Dish)
	dish.Id = 10
	helper := int(50)
	dish.CategoryID = &helper

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(dish.CategoryID, dish.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))

		res := srvc.UpdateDishCategory(ctx, dish)
		require.Nil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("Id not found")

		mock.ExpectExec(query).
			WithArgs(dish.Name, dish.Id).
			WillReturnError(expectedErr)

		err := srvc.UpdateDishCategory(ctx, dish)
		require.Error(t, err)
	})
}

func TestDishService_UpdateDishPrice(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("UPDATE pratos SET preco = ? WHERE id = ?")

	dish := new(entity.Dish)
	dish.Id = 5
	dish.Price = 35

	t.Run("success", func(t *testing.T) {

		mock.ExpectExec(query).
			WithArgs(dish.Price, dish.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))

		res := srvc.UpdateDishPrice(ctx, dish)
		require.Nil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("Id not found")

		mock.ExpectExec(query).
			WithArgs(dish.Price, dish.Id).
			WillReturnError(expectedErr)

		err := srvc.UpdateDishPrice(ctx, dish)
		require.Error(t, err)
	})
}

func TestDishService_UpdateDishPrepTime(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("UPDATE pratos SET tempo_de_preparo = ? WHERE id = ?")

	dish := new(entity.Dish)
	dish.Id = 1
	dish.CookTime = 60

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(dish.CookTime, dish.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))

		res := srvc.UpdateDishPrepTime(ctx, dish)
		require.NoError(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("Id not found")

		mock.ExpectExec(query).
			WithArgs(dish.CookTime, dish.Id).
			WillReturnError(expectedErr)

		err := srvc.UpdateDishPrepTime(ctx, dish)
		require.Error(t, err)
	})
}

func TestDishService_CreateDish(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewDishService(sqlxDB)
	query := regexp.QuoteMeta("INSERT INTO pratos (id, id_categoria, nome, preco, tempo_de_preparo) VALUES (?, ?, ?, ?, ?)")

	mockDish := new(entity.Dish)
	mockDish.Id = 1
	mockDish.CookTime = 1
	mockDish.Name = "banana"
	mockDish.Price = 1
	helper := int(1)
	mockDish.CategoryID = &helper

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(mockDish.Id, mockDish.CategoryID, mockDish.Name, mockDish.Price, mockDish.CookTime).
			WillReturnResult(sqlmock.NewResult(1, 1))

		res := srvc.CreateDish(ctx, mockDish)
		require.Nil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WillReturnError(expectedErr)

		err := srvc.CreateDish(ctx, mockDish)
		require.Error(t, err)
	})
}
