package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

func (r *queryResolver) Article(ctx context.Context) (*model.ArticleQuery, error) {
	return &model.ArticleQuery{}, nil
}

// ArticleQuery returns runtime.ArticleQueryResolver implementation.
func (r *Resolver) ArticleQuery() runtime.ArticleQueryResolver { return &articleQueryResolver{r} }

type articleQueryResolver struct{ *Resolver }
