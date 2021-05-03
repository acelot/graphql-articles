package app

import (
	"github.com/acelot/articles/internal/feature/article"
	"github.com/acelot/articles/internal/feature/articleblock"
	"github.com/acelot/articles/internal/feature/articletag"
	"github.com/acelot/articles/internal/feature/image"
	"github.com/acelot/articles/internal/feature/project"
	"github.com/acelot/articles/internal/feature/tag"
)

type Repositories struct {
	ArticleRepository      *article.Repository
	ArticleBlockRepository *articleblock.Repository
	ArticleTagRepository   *articletag.Repository
	ImageRepository        *image.Repository
	ProjectRepository      *project.Repository
	TagRepository          *tag.Repository
}

func NewRepositories(databases *Databases) *Repositories {
	return &Repositories{
		ArticleRepository:      article.NewRepository(databases.PrimaryDB, databases.SecondaryDB),
		ArticleBlockRepository: articleblock.NewRepository(databases.PrimaryDB, databases.SecondaryDB),
		ArticleTagRepository:   articletag.NewRepository(databases.PrimaryDB, databases.SecondaryDB),
		ImageRepository:        image.NewRepository(databases.PrimaryDB, databases.SecondaryDB),
		ProjectRepository:      project.NewRepository(databases.PrimaryDB, databases.SecondaryDB),
		TagRepository:          tag.NewRepository(databases.PrimaryDB, databases.SecondaryDB),
	}
}
