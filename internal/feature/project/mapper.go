package project

import (
	"github.com/acelot/articles/internal/gql/model"
)

func MapOneToGqlModel(project Project) *model.Project {
	return &model.Project{
		CreatedAt:  project.CreatedAt,
		DeletedAt:  project.DeletedAt,
		ID:         project.ID,
		ModifiedAt: project.ModifiedAt,
		Name:       project.Name,
		Version:    project.Version,
	}
}

func MapManyToGqlModels(projects []Project) []*model.Project {
	items := make([]*model.Project, len(projects))

	for i, entity := range projects {
		items[i] = MapOneToGqlModel(entity)
	}

	return items
}