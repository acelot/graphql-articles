package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/acelot/articles/internal/feature/articleblock"
	"github.com/acelot/articles/internal/feature/articletag"
	"github.com/acelot/articles/internal/gql/model"
	"go.uber.org/zap"
)

func (r *articleTagMutationResolver) Move(ctx context.Context, obj *model.ArticleTagMutation, input model.ArticleTagMoveInput) (model.ArticleTagMoveResult, error) {
	movedArticleTag, err := r.env.Services.ArticleTagService.MoveArticleTag(input)

	if errors.Is(err, articletag.NotFoundError) {
		return model.ArticleTagNotFoundProblem{Message: "article tag not found"}, nil
	}

	if errors.Is(err, articletag.InvalidSortRankError) {
		return model.InvalidSortRankProblem{Message: "outdated sort rankings"}, nil
	}

	if errors.Is(err, articleblock.VersionMismatchError) {
		return NewVersionMismatchProblem(), nil
	}

	if err != nil {
		r.env.Logger.Error("articletag.Service.MoveArticleTag", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ArticleTagMoveOk{
		SortRank: movedArticleTag.SortRank,
		Version:  movedArticleTag.Version,
	}, nil
}
