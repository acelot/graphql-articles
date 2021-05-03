package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

func (r *queryResolver) Image(ctx context.Context) (*model.ImageQuery, error) {
	return &model.ImageQuery{}, nil
}

// ImageQuery returns runtime.ImageQueryResolver implementation.
func (r *Resolver) ImageQuery() runtime.ImageQueryResolver { return &imageQueryResolver{r} }

type imageQueryResolver struct{ *Resolver }
