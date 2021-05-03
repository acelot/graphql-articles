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

func (r *articleBlockMutationResolver) Move(ctx context.Context, obj *model.ArticleBlockMutation, input model.ArticleBlockMoveInput) (model.ArticleBlockMoveResult, error) {
	movedArticleBlock, err := r.env.Services.ArticleBlockService.MoveArticleBlock(input)

	if errors.Is(err, articleblock.ArticleNotFoundError) {
		return model.ArticleBlockNotFoundProblem{Message: "article block not found"}, nil
	}

	if errors.Is(err, articleblock.InvalidSortRankError) {
		return model.InvalidSortRankProblem{Message: "outdated sort rankings"}, nil
	}

	if errors.Is(err, articleblock.VersionMismatchError) {
		return NewVersionMismatchProblem(), nil
	}

	if err != nil {
		r.env.Logger.Error("articleblock.Service.MoveArticleBlock", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ArticleBlockMoveOk{
		SortRank: movedArticleBlock.SortRank,
	}, nil
}
