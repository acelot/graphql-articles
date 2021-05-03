package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/acelot/articles/internal/feature/articleblock"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
	"go.uber.org/zap"
)

func (r *articleBlockFindListResolver) TotalCount(ctx context.Context, obj *model.ArticleBlockFindList, estimate uint) (model.TotalCountResolvingResult, error) {
	filter := graphql.GetFieldContext(ctx).Parent.Args["filter"].(*model.ArticleBlockFindFilterInput)

	count, err := r.env.Services.ArticleBlockService.CountArticleBlocks(filter, estimate)
	if err != nil {
		r.env.Logger.Error("articleblock.Service.CountArticleBlocks", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.TotalCount{
		Value: count,
	}, nil
}

func (r *articleBlockQueryResolver) Find(ctx context.Context, obj *model.ArticleBlockQuery, filter *model.ArticleBlockFindFilterInput, sort model.ArticleBlockFindSortEnum, pageSize uint, pageNumber uint) (model.ArticleBlockFindResult, error) {
	articleBlocks, err := r.env.Services.ArticleBlockService.FindArticleBlocks(filter, sort, pageSize, pageNumber)
	if err != nil {
		r.env.Logger.Error("articleblock.Service.FindArticleBlocks", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	mapped, err := articleblock.MapManyToGqlModels(articleBlocks)
	if err != nil {
		r.env.Logger.Error("articleblock.MapManyToGqlModels", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ArticleBlockFindList{
		Items: mapped,
	}, nil
}

// ArticleBlockFindList returns runtime.ArticleBlockFindListResolver implementation.
func (r *Resolver) ArticleBlockFindList() runtime.ArticleBlockFindListResolver {
	return &articleBlockFindListResolver{r}
}

type articleBlockFindListResolver struct{ *Resolver }
