package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/acelot/articles/internal/feature/image"
	"github.com/acelot/articles/internal/gql/model"
	"go.uber.org/zap"
)

func (r *imageMutationResolver) Upload(ctx context.Context, obj *model.ImageMutation, input model.ImageUploadInput) (model.ImageUploadResult, error) {
	if input.File == nil {
		return nil, errors.New("file not uploaded")
	}

	createdImage, err := r.env.Services.ImageService.UploadImage(input.File)

	if errors.Is(err, image.NotRecognizedError) {
		return model.ImageNotRecognizedProblem{Message: "image not recognized"}, nil
	}

	if errors.Is(err, image.NotSupportedTypeError) {
		return model.ImageNotSupportedTypeProblem{
			Message: "image type not supported; supported formats: jpeg, png, webp, avif",
		}, nil
	}

	if err != nil {
		r.env.Logger.Error("image.Service.UploadImage", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	gqlImage, err := image.MapOneToGqlModel(*createdImage)
	if err != nil {
		r.env.Logger.Error("image.MapOneToGqlModel", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ImageUploadOk{
		Image: gqlImage,
	}, nil
}
