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

func (r *articleBlockMutationResolver) Create(ctx context.Context, obj *model.ArticleBlockMutation, input model.ArticleBlockCreateInput) (model.ArticleBlockCreateResult, error) {
	createdArticleBlock, err := r.env.Services.ArticleBlockService.CreateArticleBlock(input)

	if errors.Is(err, articleblock.ArticleNotFoundError) {
		return model.ArticleNotFoundProblem{Message: "article not found"}, nil
	}

	if errors.Is(err, articleblock.InvalidSortRankError) {
		return model.InvalidSortRankProblem{Message: "outdated sort rankings"}, nil
	}

	if err != nil {
		r.env.Logger.Error("articleblock.Service.CreateArticleBlock", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	gqlArticleBlock, err := articleblock.MapOneToGqlModel(*createdArticleBlock)
	if err != nil {
		r.env.Logger.Error("articleblock.MapOneToGqlModel", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ArticleBlockCreateOk{
		ArticleBlock: gqlArticleBlock,
	}, nil
}
