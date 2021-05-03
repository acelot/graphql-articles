package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

func (r *mutationResolver) Project(ctx context.Context) (*model.ProjectMutation, error) {
	return &model.ProjectMutation{}, nil
}

// ProjectMutation returns runtime.ProjectMutationResolver implementation.
func (r *Resolver) ProjectMutation() runtime.ProjectMutationResolver {
	return &projectMutationResolver{r}
}

type projectMutationResolver struct{ *Resolver }
