package handler

import (
	"github.com/acelot/articles/internal/app"
	"github.com/acelot/articles/internal/http/middleware"
	"net/http"
)

func NewAppHandler(env *app.Env) *http.ServeMux {
	dataLoaderInjector := middleware.NewDataLoadersInjector(env)

	mux := http.NewServeMux()
	mux.Handle("/api", dataLoaderInjector(NewGraphQLHandler(env)))

	return mux
}
