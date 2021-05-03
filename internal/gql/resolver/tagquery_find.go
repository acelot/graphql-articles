package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/acelot/articles/internal/feature/tag"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
	"go.uber.org/zap"
)

func (r *tagFindListResolver) TotalCount(ctx context.Context, obj *model.TagFindList, estimate uint) (model.TotalCountResolvingResult, error) {
	filter := graphql.GetFieldContext(ctx).Parent.Args["filter"].(*model.TagFindFilterInput)

	count, err := r.env.Services.TagService.CountTags(filter, estimate)
	if err != nil {
		r.env.Logger.Error("tag.Service.CountTags", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.TotalCount{
		Value: count,
	}, nil
}

func (r *tagQueryResolver) Find(ctx context.Context, obj *model.TagQuery, filter *model.TagFindFilterInput, sort model.TagFindSortEnum, pageSize uint, pageNumber uint) (model.TagFindResult, error) {
	tags, err := r.env.Services.TagService.FindTags(filter, sort, pageSize, pageNumber)
	if err != nil {
		r.env.Logger.Error("tag.Service.FindTags", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.TagFindList{
		Items: tag.MapManyToGqlModels(tags),
	}, nil
}

// TagFindList returns runtime.TagFindListResolver implementation.
func (r *Resolver) TagFindList() runtime.TagFindListResolver { return &tagFindListResolver{r} }

type tagFindListResolver struct{ *Resolver }
