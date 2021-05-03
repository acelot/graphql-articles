package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

func (r *queryResolver) Tag(ctx context.Context) (*model.TagQuery, error) {
	return &model.TagQuery{}, nil
}

// TagQuery returns runtime.TagQueryResolver implementation.
func (r *Resolver) TagQuery() runtime.TagQueryResolver { return &tagQueryResolver{r} }

type tagQueryResolver struct{ *Resolver }
