package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

func (r *mutationResolver) Article(ctx context.Context) (*model.ArticleMutation, error) {
	return &model.ArticleMutation{}, nil
}

// ArticleMutation returns runtime.ArticleMutationResolver implementation.
func (r *Resolver) ArticleMutation() runtime.ArticleMutationResolver {
	return &articleMutationResolver{r}
}

type articleMutationResolver struct{ *Resolver }
