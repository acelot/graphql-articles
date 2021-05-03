package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

func (r *mutationResolver) Tag(ctx context.Context) (*model.TagMutation, error) {
	return &model.TagMutation{}, nil
}

// TagMutation returns runtime.TagMutationResolver implementation.
func (r *Resolver) TagMutation() runtime.TagMutationResolver { return &tagMutationResolver{r} }

type tagMutationResolver struct{ *Resolver }
