package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

func (r *imageResolver) Assets(ctx context.Context, obj *model.Image) ([]*model.ImageAsset, error) {
	return []*model.ImageAsset{}, nil
}

func (r *imageResolver) Download(ctx context.Context, obj *model.Image) (model.ImageDownloadResolvingResult, error) {
	tempURL, err := r.env.Services.ImageService.GetImageDownloadURL(obj.ID)
	if err != nil {
		return NewInternalErrorProblem(), nil
	}

	return model.ImageDownload{URL: *tempURL}, nil
}

// Image returns runtime.ImageResolver implementation.
func (r *Resolver) Image() runtime.ImageResolver { return &imageResolver{r} }

type imageResolver struct{ *Resolver }
