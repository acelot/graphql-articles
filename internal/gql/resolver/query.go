package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"github.com/acelot/articles/internal/gql/runtime"
)

// Query returns runtime.QueryResolver implementation.
func (r *Resolver) Query() runtime.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
