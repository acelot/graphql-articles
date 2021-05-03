package article

import (
	"github.com/acelot/articles/internal/gql/model"
)

func MapOneToGqlModel(article Article) *model.Article {
	return &model.Article{
		CoverImageID: article.CoverImageID,
		CreatedAt:    article.CreatedAt,
		DeletedAt:    article.DeletedAt,
		ID:           article.ID,
		ModifiedAt:   article.ModifiedAt,
		ProjectID:    article.ProjectID,
		Title:        article.Title,
		Version:      article.Version,
	}
}

func MapManyToGqlModels(articles []Article) []*model.Article {
	items := make([]*model.Article, len(articles))

	for i, entity := range articles {
		items[i] = MapOneToGqlModel(entity)
	}

	return items
}
