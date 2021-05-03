package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/acelot/articles/internal/feature/article"
	"github.com/acelot/articles/internal/gql/model"
	"go.uber.org/zap"
)

func (r *articleMutationResolver) Create(ctx context.Context, obj *model.ArticleMutation, input model.ArticleCreateInput) (model.ArticleCreateResult, error) {
	createdArticle, err := r.env.Services.ArticleService.CreateArticle(input.ProjectID)

	if errors.Is(err, article.ProjectIDConstraintError) {
		return model.ProjectNotFoundProblem{Message: "project not found"}, nil
	}

	if err != nil {
		r.env.Logger.Error("article.Service.CreateArticle", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ArticleCreateOk{
		Article: article.MapOneToGqlModel(*createdArticle),
	}, nil
}
