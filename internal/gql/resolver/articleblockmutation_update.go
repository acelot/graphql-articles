package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/acelot/articles/internal/feature/articleblock"
	"github.com/acelot/articles/internal/gql/model"
	"go.uber.org/zap"
)

func (r *articleBlockMutationResolver) Update(ctx context.Context, obj *model.ArticleBlockMutation, input model.ArticleBlockUpdateInput) (model.ArticleBlockUpdateResult, error) {
	updatedArticleBlock, err := r.env.Services.ArticleBlockService.UpdateArticleBlock(input)

	if errors.Is(err, articleblock.ArticleNotFoundError) {
		return model.ArticleBlockNotFoundProblem{Message: "article block not found"}, nil
	}

	if errors.Is(err, articleblock.TypeMismatchError) {
		return model.ArticleBlockTypeMismatchProblem{
			Message: "data input type of the article block doesn't match the block type of the updated block",
		}, nil
	}

	if errors.Is(err, articleblock.VersionMismatchError) {
		return NewVersionMismatchProblem(), nil
	}

	if err != nil {
		r.env.Logger.Error("articleblock.Service.UpdateArticleBlock", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	gqlArticleBlock, err := articleblock.MapOneToGqlModel(*updatedArticleBlock)
	if err != nil {
		r.env.Logger.Error("articleblock.MapOneToGqlModel", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ArticleBlockUpdateOk{
		ArticleBlock: gqlArticleBlock,
	}, nil
}
