package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/acelot/articles/internal/feature/project"
	"github.com/acelot/articles/internal/gql/model"
	"go.uber.org/zap"
)

func (r *projectMutationResolver) Create(ctx context.Context, obj *model.ProjectMutation, input model.ProjectCreateInput) (model.ProjectCreateResult, error) {
	createdProject, err := r.env.Services.ProjectService.CreateProject(input)

	if errors.Is(err, project.ProjectNameConstraintError) {
		return model.ProjectAlreadyExistsProblem{Message: "project already exists"}, nil
	}

	if err != nil {
		r.env.Logger.Error("project.Service.CreateProject", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ProjectCreateOk{
		Project: project.MapOneToGqlModel(*createdProject),
	}, nil
}
