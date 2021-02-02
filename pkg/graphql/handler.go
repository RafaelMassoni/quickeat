package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"

	"quickeat/pkg/graphql/gqlgen"
	graphql "quickeat/pkg/graphql/resolvers"
)

func NewHandler(db *sqlx.DB) http.Handler {
	r := chi.NewRouter()
	gqlgenConfig := gqlgen.Config{
		Resolvers: graphql.NewResolverRoot(db),
	}

	svc := handler.NewDefaultServer(gqlgen.NewExecutableSchema(gqlgenConfig))

	r.Handle("/", svc)
	r.Handle("/explorer", playground.Handler("Explorer", "/"))
	return r
}
