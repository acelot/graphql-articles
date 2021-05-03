package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

func (r *mutationResolver) ArticleTag(ctx context.Context) (*model.ArticleTagMutation, error) {
	return &model.ArticleTagMutation{}, nil
}

// ArticleTagMutation returns runtime.ArticleTagMutationResolver implementation.
func (r *Resolver) ArticleTagMutation() runtime.ArticleTagMutationResolver {
	return &articleTagMutationResolver{r}
}

type articleTagMutationResolver struct{ *Resolver }
