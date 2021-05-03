package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/acelot/articles/internal/feature/articletag"
	"github.com/acelot/articles/internal/gql/model"
	"go.uber.org/zap"
)

func (r *articleTagMutationResolver) Create(ctx context.Context, obj *model.ArticleTagMutation, input model.ArticleTagCreateInput) (model.ArticleTagCreateResult, error) {
	createdArticleTag, err := r.env.Services.ArticleTagService.CreateArticleTag(input)

	if errors.Is(err, articletag.ArticleNotFoundError) {
		return model.ArticleNotFoundProblem{Message: "article not found"}, nil
	}

	if errors.Is(err, articletag.TagNotFoundError) {
		return model.TagNotFoundProblem{Message: "tag not found"}, nil
	}

	if errors.Is(err, articletag.DuplicateTagError) {
		return model.ArticleTagAlreadyExistsProblem{Message: "tag already exists in the article"}, nil
	}

	if errors.Is(err, articletag.InvalidSortRankError) {
		return model.InvalidSortRankProblem{Message: "outdated sort rankings"}, nil
	}

	if err != nil {
		r.env.Logger.Error("articletag.Service.CreateArticleTag", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ArticleTagCreateOk{
		ArticleTag: articletag.MapOneToGqlModel(*createdArticleTag),
	}, nil
}
