package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/acelot/articles/internal/app"
	"github.com/acelot/articles/internal/feature/tag"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
	"go.uber.org/zap"
)

func (r *articleTagResolver) Tag(ctx context.Context, obj *model.ArticleTag) (model.TagResolvingResult, error) {
	dataLoader := ctx.Value(app.DataLoadersContextKey).(*app.DataLoaders).TagLoaderByID

	proj, err := dataLoader.Load(tag.LoaderByIDKey{ID: obj.TagID})
	if err != nil {
		r.env.Logger.Error("tag.LoaderByID.Load", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return tag.MapOneToGqlModel(*proj), nil
}

// ArticleTag returns runtime.ArticleTagResolver implementation.
func (r *Resolver) ArticleTag() runtime.ArticleTagResolver { return &articleTagResolver{r} }

type articleTagResolver struct{ *Resolver }
