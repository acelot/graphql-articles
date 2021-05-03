package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

func (r *queryResolver) Project(ctx context.Context) (*model.ProjectQuery, error) {
	return &model.ProjectQuery{}, nil
}

// ProjectQuery returns runtime.ProjectQueryResolver implementation.
func (r *Resolver) ProjectQuery() runtime.ProjectQueryResolver { return &projectQueryResolver{r} }

type projectQueryResolver struct{ *Resolver }
