package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/app"
	"github.com/acelot/articles/internal/feature/image"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
	"go.uber.org/zap"
)

func (r *articleBlockImageDataResolver) Image(ctx context.Context, obj *model.ArticleBlockImageData) (model.ImageResolvingResult, error) {
	if obj.ImageID == nil {
		return model.ImageNotFoundProblem{Message: "image not set"}, nil
	}

	dataLoader := ctx.Value(app.DataLoadersContextKey).(*app.DataLoaders).ImageLoaderByID

	img, err := dataLoader.Load(image.LoaderByIDKey{ID: *obj.ImageID})
	if err != nil {
		r.env.Logger.Error("image.LoaderByID.Load", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	if img == nil {
		return model.ImageNotFoundProblem{Message: "image not found"}, nil
	}

	gqlImage, err := image.MapOneToGqlModel(*img)
	if err != nil {
		r.env.Logger.Error("image.MapOneToGqlModel", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return *gqlImage, nil
}

// ArticleBlockImageData returns runtime.ArticleBlockImageDataResolver implementation.
func (r *Resolver) ArticleBlockImageData() runtime.ArticleBlockImageDataResolver {
	return &articleBlockImageDataResolver{r}
}

type articleBlockImageDataResolver struct{ *Resolver }
