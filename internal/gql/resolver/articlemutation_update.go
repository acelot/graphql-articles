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

func (r *articleMutationResolver) Update(ctx context.Context, obj *model.ArticleMutation, input model.ArticleUpdateInput) (model.ArticleUpdateResult, error) {
	updatedArticle, err := r.env.Services.ArticleService.UpdateArticle(input)

	if errors.Is(err, article.NotFoundError) {
		return model.ArticleNotFoundProblem{Message: "article not found"}, nil
	}

	if errors.Is(err, article.VersionMismatchError) {
		return NewVersionMismatchProblem(), nil
	}

	if err != nil {
		r.env.Logger.Error("article.Service.UpdateArticle", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ArticleUpdateOk{
		Article: article.MapOneToGqlModel(*updatedArticle),
	}, nil
}
