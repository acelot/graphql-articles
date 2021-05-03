package articletag

import (
	"github.com/acelot/articles/internal/gql/model"
)

func MapOneToGqlModel(articleTag ArticleTag) *model.ArticleTag {
	return &model.ArticleTag{
		ArticleID:  articleTag.ArticleID,
		CreatedAt:  articleTag.CreatedAt,
		ID:         articleTag.ID,
		ModifiedAt: articleTag.ModifiedAt,
		SortRank:   articleTag.SortRank,
		TagID:      articleTag.TagID,
		Version:    articleTag.Version,
	}
}

func MapManyToGqlModels(articleTags []ArticleTag) []*model.ArticleTag {
	items := make([]*model.ArticleTag, len(articleTags))

	for i, entity := range articleTags {
		items[i] = MapOneToGqlModel(entity)
	}

	return items
}
