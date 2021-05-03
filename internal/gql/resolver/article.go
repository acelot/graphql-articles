package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/app"
	"github.com/acelot/articles/internal/feature/articleblock"
	"github.com/acelot/articles/internal/feature/articletag"
	"github.com/acelot/articles/internal/feature/image"
	"github.com/acelot/articles/internal/feature/project"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
	"go.uber.org/zap"
)

func (r *articleResolver) Content(ctx context.Context, obj *model.Article) (model.ArticleContentResolvingResult, error) {
	dataLoader := ctx.Value(app.DataLoadersContextKey).(*app.DataLoaders).ArticleBlockLoaderByArticleID

	articleBlocks, err := dataLoader.Load(articleblock.LoaderByArticleIDKey{ArticleID: obj.ID})
	if err != nil {
		r.env.Logger.Error("articleblock.LoaderByArticleID.Load", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	gqlArticleBlocks, err := articleblock.MapManyToGqlModels(articleBlocks)
	if err != nil {
		r.env.Logger.Error("articleblock.MapManyToGqlModels", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ArticleContent{
		Blocks: gqlArticleBlocks,
	}, nil
}

func (r *articleResolver) CoverImage(ctx context.Context, obj *model.Article) (model.ImageResolvingResult, error) {
	dataLoader := ctx.Value(app.DataLoadersContextKey).(*app.DataLoaders).ImageLoaderByID

	if obj.CoverImageID == nil {
		return model.ImageNotFoundProblem{
			Message: "no cover image",
		}, nil
	}

	img, err := dataLoader.Load(image.LoaderByIDKey{ID: *obj.CoverImageID})
	if err != nil {
		r.env.Logger.Error("image.LoaderByID.Load", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	if img == nil {
		return model.ImageNotFoundProblem{
			Message: "cover image not found",
		}, nil
	}

	gqlImage, err := image.MapOneToGqlModel(*img)
	if err != nil {
		r.env.Logger.Error("image.MapOneToGqlModel", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return *gqlImage, nil
}

func (r *articleResolver) Project(ctx context.Context, obj *model.Article) (model.ProjectResolvingResult, error) {
	dataLoader := ctx.Value(app.DataLoadersContextKey).(*app.DataLoaders).ProjectLoaderByID

	proj, err := dataLoader.Load(project.LoaderByIDKey{ID: obj.ProjectID})
	if err != nil {
		r.env.Logger.Error("project.LoaderByID.Load", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	if proj == nil {
		return model.ProjectNotFoundProblem{
			Message: "project not found",
		}, nil
	}

	return project.MapOneToGqlModel(*proj), nil
}

func (r *articleResolver) Tags(ctx context.Context, obj *model.Article) (model.ArticleTagsResolvingResult, error) {
	dataLoader := ctx.Value(app.DataLoadersContextKey).(*app.DataLoaders).ArticleTagLoaderByArticleID

	articleTags, err := dataLoader.Load(articletag.LoaderByArticleIDKey{ArticleID: obj.ID})
	if err != nil {
		r.env.Logger.Error("articletag.LoaderByArticleID.Load", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.ArticleTagList{
		Items: articletag.MapManyToGqlModels(articleTags),
	}, nil
}

// Article returns runtime.ArticleResolver implementation.
func (r *Resolver) Article() runtime.ArticleResolver { return &articleResolver{r} }

type articleResolver struct{ *Resolver }
