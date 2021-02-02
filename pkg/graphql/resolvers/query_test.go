package graphql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryResolver_Dish(t *testing.T) {
	ctx := context.Background()
	query := NewQueryResolver()

	t.Run("success", func(t *testing.T) {
		name := "any"
		category := []string{"any"}

		res, err := query.Dish(ctx, name, category)

		require.Nil(t, err)
		require.Equal(t, res[0].Name, "Peito de Frango")
	})
}
