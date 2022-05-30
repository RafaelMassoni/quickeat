package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"testing"

	"quickeat/pkg/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestCategoryService_Get(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewCategoryService(sqlxDB)

	categoryId := 10

	query := regexp.QuoteMeta(fmt.Sprintf("SELECT * FROM categorias WHERE id = %d", categoryId))

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(query).
			WillReturnRows(sqlmock.NewRows([]string{"id", "nome"}).
				AddRow(categoryId, "massas"))

		res, err := srvc.Get(ctx, &categoryId)
		require.NoError(t, err)
		require.Equal(t, categoryId, res[0].Id)
		require.Equal(t, "massas", res[0].Name)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WithArgs(categoryId).
			WillReturnError(expectedErr)

		res, err := srvc.Get(ctx, &categoryId)
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestCategoryService_GetByDish(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewCategoryService(sqlxDB)

	query := regexp.QuoteMeta(
		"SELECT c.id, c.nome " +
			"FROM categorias as c " +
			"INNER JOIN pratos as p ON c.id = p.id_categoria " +
			"WHERE p.id = ?")

	t.Run("success", func(t *testing.T) {
		dishId := 200
		categoryId := 10
		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "nome"}).
				AddRow(categoryId, "massas"))

		res, err := srvc.GetByDish(ctx, dishId)
		require.NoError(t, err)
		require.Equal(t, categoryId, res.Id)
		require.Equal(t, "massas", res.Name)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")
		dishId := 200

		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		res, err := srvc.GetByDish(ctx, dishId)
		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("failed: no rows", func(t *testing.T) {
		expectedErr := sql.ErrNoRows
		dishId := 10
		mock.ExpectQuery(query).
			WithArgs(dishId).
			WillReturnError(expectedErr)

		res, err := srvc.GetByDish(ctx, dishId)
		require.Nil(t, err)
		require.Nil(t, res)
	})
}

func TestCategoryService_Create(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewCategoryService(sqlxDB)

	mockCategory := new(entity.Category)
	mockCategory.Id = 10
	mockCategory.Name = "japonesa"

	query := regexp.QuoteMeta("INSERT INTO categorias (id, nome) VALUES (?, ?)")

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(mockCategory.Id, mockCategory.Name).
			WillReturnResult(sqlmock.NewResult(1, 1))

		res := srvc.Create(ctx, mockCategory)
		require.Nil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WillReturnError(expectedErr)

		res := srvc.Create(ctx, mockCategory)
		require.Error(t, res)
	})
}

func TestCategoryService_Update(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewCategoryService(sqlxDB)

	mockCategory := new(entity.Category)
	mockCategory.Id = 10
	mockCategory.Name = "japonesa"

	query := regexp.QuoteMeta("UPDATE categorias SET nome = ? WHERE id = ?")

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(mockCategory.Name, mockCategory.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))

		res := srvc.Update(ctx, mockCategory)
		require.Nil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WillReturnError(expectedErr)

		res := srvc.Update(ctx, mockCategory)
		require.Error(t, res)
	})
}

func TestCategoryService_DeleteCategoryByName(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewCategoryService(sqlxDB)
	query := regexp.QuoteMeta("DELETE FROM categorias WHERE nome = ?")

	t.Run("success", func(t *testing.T) {
		categoryName := "massas"

		mock.ExpectQuery(query).
			WithArgs(categoryName).
			WillReturnRows(sqlmock.NewRows([]string{"id", "nome"}).
				AddRow(1, categoryName))

		res := srvc.DeleteCategoryByName(ctx, categoryName)
		require.Nil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		categoryName := "massas"
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WithArgs(categoryName).
			WillReturnError(expectedErr)

		err := srvc.DeleteCategoryByName(ctx, categoryName)
		require.Error(t, err)
	})
}

func TestCategoryService_DeleteCategoryById(t *testing.T) {
	ctx := context.Background()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	srvc := NewCategoryService(sqlxDB)
	query := regexp.QuoteMeta("DELETE FROM categorias WHERE id = ?")

	t.Run("success", func(t *testing.T) {
		categoryId := 1

		mock.ExpectQuery(query).
			WithArgs(categoryId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "nome"}).
				AddRow(categoryId, "massas"))

		res := srvc.DeleteCategoryById(ctx, categoryId)
		require.Nil(t, res)
	})

	t.Run("failed", func(t *testing.T) {
		categoryId := 1
		expectedErr := errors.New("db failed")

		mock.ExpectQuery(query).
			WithArgs(categoryId).
			WillReturnError(expectedErr)

		err := srvc.DeleteCategoryById(ctx, categoryId)
		require.Error(t, err)
	})
}
