package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/acelot/articles/internal/feature/project"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
	"go.uber.org/zap"
)

func (r *projectFindListResolver) TotalCount(ctx context.Context, obj *model.ProjectFindList) (model.TotalCountResolvingResult, error) {
	filter := graphql.GetFieldContext(ctx).Parent.Args["filter"].(*model.ProjectFindFilterInput)

	count, err := r.env.Services.ProjectService.CountProjects(filter)
	if err != nil {
		r.env.Logger.Error("project.Service.CountProjects", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.TotalCount{
		Value: count,
	}, nil
}

func (r *projectQueryResolver) Find(ctx context.Context, obj *model.ProjectQuery, filter *model.ProjectFindFilterInput, sort model.ProjectFindSortEnum, pageSize uint, pageNumber uint) (model.ProjectFindResult, error) {
	projects, err := r.env.Services.ProjectService.FindProjects(filter, sort, pageSize, pageNumber)
	if err != nil {
		r.env.Logger.Error("project.Service.FindProjects", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ProjectFindList{
		Items: project.MapManyToGqlModels(projects),
	}, nil
}

// ProjectFindList returns runtime.ProjectFindListResolver implementation.
func (r *Resolver) ProjectFindList() runtime.ProjectFindListResolver {
	return &projectFindListResolver{r}
}

type projectFindListResolver struct{ *Resolver }
