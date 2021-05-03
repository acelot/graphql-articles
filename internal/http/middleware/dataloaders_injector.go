package middleware

import (
	"context"
	"github.com/acelot/articles/internal/app"
	"net/http"
)

func NewDataLoadersInjector(env *app.Env) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(
				r.Context(),
				app.DataLoadersContextKey,
				app.NewDataLoaders(env.Repositories),
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
