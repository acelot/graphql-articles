package app

import (
	"github.com/acelot/articles/internal/feature/article"
	"github.com/acelot/articles/internal/feature/articleblock"
	"github.com/acelot/articles/internal/feature/articletag"
	"github.com/acelot/articles/internal/feature/image"
	"github.com/acelot/articles/internal/feature/project"
	"github.com/acelot/articles/internal/feature/tag"
)

type Services struct {
	ArticleService      *article.Service
	ArticleBlockService *articleblock.Service
	ArticleTagService   *articletag.Service
	ImageService        *image.Service
	TagService          *tag.Service
	ProjectService      *project.Service
}

func NewServices(repos *Repositories, storages *Storages) *Services {
	return &Services{
		ArticleService:      article.NewService(repos.ArticleRepository),
		ArticleBlockService: articleblock.NewService(repos.ArticleBlockRepository),
		ArticleTagService:   articletag.NewService(repos.ArticleTagRepository),
		ImageService:        image.NewService(repos.ImageRepository, storages.ImageStorage),
		TagService:          tag.NewService(repos.TagRepository),
		ProjectService:      project.NewService(repos.ProjectRepository),
	}
}
