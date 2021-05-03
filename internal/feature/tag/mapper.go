package tag

import (
	"github.com/acelot/articles/internal/gql/model"
)

func MapOneToGqlModel(tag Tag) *model.Tag {
	return &model.Tag{
		CreatedAt:  tag.CreatedAt,
		DeletedAt:  tag.DeletedAt,
		ID:         tag.ID,
		ModifiedAt: tag.ModifiedAt,
		Name:       tag.Name,
		Version:    tag.Version,
	}
}

func MapManyToGqlModels(projects []Tag) []*model.Tag {
	items := make([]*model.Tag, len(projects))

	for i, entity := range projects {
		items[i] = MapOneToGqlModel(entity)
	}

	return items
}