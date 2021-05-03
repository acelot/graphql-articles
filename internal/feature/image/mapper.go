package image

import (
	"github.com/acelot/articles/internal/gql/model"
)

func MapOneToGqlModel(image Image) (*model.Image, error) {
	return &model.Image{
		CreatedAt:  image.CreatedAt,
		DeletedAt:  image.DeletedAt,
		Type:       image.Type,
		Height:     uint(image.Height),
		ID:         image.ID,
		ModifiedAt: image.ModifiedAt,
		Version:    image.Version,
		Width:      uint(image.Width),
	}, nil
}

func MapManyToGqlModels(images []Image) ([]*model.Image, error) {
	items := make([]*model.Image, len(images))

	for i, entity := range images {
		m, err := MapOneToGqlModel(entity)
		if err != nil {
			return nil, err
		}

		items[i] = m
	}

	return items, nil
}