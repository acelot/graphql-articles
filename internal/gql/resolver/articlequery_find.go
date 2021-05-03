package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/acelot/articles/internal/feature/article"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
	"go.uber.org/zap"
)

func (r *articleFindListResolver) TotalCount(ctx context.Context, obj *model.ArticleFindList, estimate uint) (model.TotalCountResolvingResult, error) {
	filter := graphql.GetFieldContext(ctx).Parent.Args["filter"].(*model.ArticleFindFilterInput)

	count, err := r.env.Services.ArticleService.CountArticles(filter, estimate)
	if err != nil {
		r.env.Logger.Error("article.Service.CountArticles", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.TotalCount{
		Value: count,
	}, nil
}

func (r *articleQueryResolver) Find(ctx context.Context, obj *model.ArticleQuery, filter *model.ArticleFindFilterInput, sort model.ArticleFindSortEnum, pageSize uint, pageNumber uint) (model.ArticleFindResult, error) {
	articles, err := r.env.Services.ArticleService.FindArticles(filter, sort, pageSize, pageNumber)
	if err != nil {
		r.env.Logger.Error("article.Service.FindArticles", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ArticleFindList{
		Items: article.MapManyToGqlModels(articles),
	}, nil
}

// ArticleFindList returns runtime.ArticleFindListResolver implementation.
func (r *Resolver) ArticleFindList() runtime.ArticleFindListResolver {
	return &articleFindListResolver{r}
}

type articleFindListResolver struct{ *Resolver }
