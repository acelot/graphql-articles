package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

func (r *mutationResolver) ArticleBlock(ctx context.Context) (*model.ArticleBlockMutation, error) {
	return &model.ArticleBlockMutation{}, nil
}

// ArticleBlockMutation returns runtime.ArticleBlockMutationResolver implementation.
func (r *Resolver) ArticleBlockMutation() runtime.ArticleBlockMutationResolver {
	return &articleBlockMutationResolver{r}
}

type articleBlockMutationResolver struct{ *Resolver }
