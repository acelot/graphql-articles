package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/acelot/articles/internal/feature/project"
	"github.com/acelot/articles/internal/feature/tag"
	"github.com/acelot/articles/internal/gql/model"
	"go.uber.org/zap"
)

func (r *projectMutationResolver) Update(ctx context.Context, obj *model.ProjectMutation, input model.ProjectUpdateInput) (model.ProjectUpdateResult, error) {
	updatedProject, err := r.env.Services.ProjectService.UpdateProject(input)

	if errors.Is(err, tag.NotFoundError) {
		return model.ProjectNotFoundProblem{Message: "project not found"}, nil
	}

	if errors.Is(err, tag.VersionMismatchError) {
		return NewVersionMismatchProblem(), nil
	}

	if err != nil {
		r.env.Logger.Error("project.Service.UpdateProject", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ProjectUpdateOk{
		Project: project.MapOneToGqlModel(*updatedProject),
	}, nil
}
