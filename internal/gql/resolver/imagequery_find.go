package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/acelot/articles/internal/feature/image"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
	"go.uber.org/zap"
)

func (r *imageFindListResolver) TotalCount(ctx context.Context, obj *model.ImageFindList, estimate uint) (model.TotalCountResolvingResult, error) {
	filter := graphql.GetFieldContext(ctx).Parent.Args["filter"].(*model.ImageFindFilterInput)

	count, err := r.env.Services.ImageService.CountImages(filter, estimate)

	if err != nil {
		r.env.Logger.Error("image.Service.CountImages", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.TotalCount{
		Value: count,
	}, nil
}

func (r *imageQueryResolver) Find(ctx context.Context, obj *model.ImageQuery, filter *model.ImageFindFilterInput, sort model.ImageFindSortEnum, pageSize uint, pageNumber uint) (model.ImageFindResult, error) {
	images, err := r.env.Services.ImageService.FindImages(filter, sort, pageSize, pageNumber)
	if err != nil {
		r.env.Logger.Error("image.Repository.Find", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	gqlImages, err := image.MapManyToGqlModels(images)
	if err != nil {
		r.env.Logger.Error("image.MapManyToGqlModels", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ImageFindList{
		Items: gqlImages,
	}, nil
}

// ImageFindList returns runtime.ImageFindListResolver implementation.
func (r *Resolver) ImageFindList() runtime.ImageFindListResolver { return &imageFindListResolver{r} }

type imageFindListResolver struct{ *Resolver }
