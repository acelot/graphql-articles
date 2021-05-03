package handler

//goland:noinspection SpellCheckingInspection
import (
	"context"
	"fmt"
	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/acelot/articles/internal/app"
	"github.com/acelot/articles/internal/gql/directive"
	"github.com/acelot/articles/internal/gql/resolver"
	"github.com/acelot/articles/internal/gql/runtime"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
	"time"
)

const websocketKeepAlivePingInterval = 5 * time.Second
const maxUploadSize = 30 * 1024 * 1024
const queryCacheLRUSize = 1000
const automaticPersistedQueryCacheLRUSize = 100
const complexityLimit = 1000

func NewGraphQLHandler(env *app.Env) *gqlhandler.Server {
	handler := gqlhandler.New(
		runtime.NewExecutableSchema(
			newSchemaConfig(env),
		),
	)

	// Transports
	handler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: websocketKeepAlivePingInterval,
	})
	handler.AddTransport(transport.Options{})
	handler.AddTransport(transport.POST{})
	handler.AddTransport(transport.MultipartForm{
		MaxUploadSize: maxUploadSize,
		MaxMemory:     maxUploadSize / 10,
	})

	// Query cache
	handler.SetQueryCache(lru.New(queryCacheLRUSize))

	// Enabling introspection
	handler.Use(extension.Introspection{})

	// APQ
	handler.Use(extension.AutomaticPersistedQuery{Cache: lru.New(automaticPersistedQueryCacheLRUSize)})

	// Complexity
	handler.Use(extension.FixedComplexityLimit(complexityLimit))

	// Unhandled errors logger
	handler.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		env.Logger.Error("unhandled error", zap.String("error", fmt.Sprintf("%v", err)))

		return gqlerror.Errorf("internal server error")
	})

	return handler
}

func newSchemaConfig(env *app.Env) runtime.Config {
	cfg := runtime.Config{Resolvers: resolver.NewResolver(env)}

	cfg.Directives.InputUnion = directive.NewInputUnionDirective()
	cfg.Directives.SortRankInput = directive.NewSortRankInputDirective()

	cfg.Complexity.ArticleQuery.Find = resolver.ArticleQueryFindComplexity
	cfg.Complexity.ArticleFindList.TotalCount = resolver.ArticleFindListTotalCountComplexity

	return cfg
}
