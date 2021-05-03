package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"github.com/acelot/articles/internal/gql/runtime"
)

// Mutation returns runtime.MutationResolver implementation.
func (r *Resolver) Mutation() runtime.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
