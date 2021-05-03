package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/acelot/articles/internal/feature/tag"
	"github.com/acelot/articles/internal/gql/model"
	"go.uber.org/zap"
)

func (r *tagMutationResolver) Create(ctx context.Context, obj *model.TagMutation, input model.TagCreateInput) (model.TagCreateResult, error) {
	createdTag, err := r.env.Services.TagService.CreateTag(input)

	if errors.Is(err, tag.TagNameConstraintError) {
		return model.TagAlreadyExistsProblem{Message: "tag already exists"}, nil
	}

	if err != nil {
		r.env.Logger.Error("tag.Service.CreateTag", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.TagCreateOk{
		Tag: tag.MapOneToGqlModel(*createdTag),
	}, nil
}
