package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

func (r *mutationResolver) Image(ctx context.Context) (*model.ImageMutation, error) {
	return &model.ImageMutation{}, nil
}

// ImageMutation returns runtime.ImageMutationResolver implementation.
func (r *Resolver) ImageMutation() runtime.ImageMutationResolver { return &imageMutationResolver{r} }

type imageMutationResolver struct{ *Resolver }
