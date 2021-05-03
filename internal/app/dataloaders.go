package app

import (
	"github.com/acelot/articles/internal/feature/article"
	"github.com/acelot/articles/internal/feature/articleblock"
	"github.com/acelot/articles/internal/feature/articletag"
	"github.com/acelot/articles/internal/feature/image"
	"github.com/acelot/articles/internal/feature/project"
	"github.com/acelot/articles/internal/feature/tag"
)

const (
	articleLoaderByIDMaxBatch             int = 100
	articleBlockLoaderByArticleIDMaxBatch int = 10
	articleTagLoaderByArticleIDMaxBatch   int = 10
	imageLoaderByIDMaxBatch               int = 100
	projectLoaderByIDMaxBatch             int = 10
	tagLoaderByIDMaxBatch                 int = 100
)

type DataLoaders struct {
	ArticleLoaderByID             *article.LoaderByID
	ArticleBlockLoaderByArticleID *articleblock.LoaderByArticleID
	ArticleTagLoaderByArticleID   *articletag.LoaderByArticleID
	ImageLoaderByID               *image.LoaderByID
	ProjectLoaderByID             *project.LoaderByID
	TagLoaderByID                 *tag.LoaderByID
}

func NewDataLoaders(repos *Repositories) *DataLoaders {
	return &DataLoaders{
		ArticleLoaderByID:             article.NewConfiguredLoaderByID(repos.ArticleRepository, articleLoaderByIDMaxBatch),
		ArticleBlockLoaderByArticleID: articleblock.NewConfiguredLoaderByArticleID(repos.ArticleBlockRepository, articleBlockLoaderByArticleIDMaxBatch),
		ArticleTagLoaderByArticleID:   articletag.NewConfiguredLoaderByArticleID(repos.ArticleTagRepository, articleTagLoaderByArticleIDMaxBatch),
		ImageLoaderByID:               image.NewConfiguredLoaderByID(repos.ImageRepository, imageLoaderByIDMaxBatch),
		ProjectLoaderByID:             project.NewConfiguredLoaderByID(repos.ProjectRepository, projectLoaderByIDMaxBatch),
		TagLoaderByID:                 tag.NewConfiguredLoaderByID(repos.TagRepository, tagLoaderByIDMaxBatch),
	}
}
