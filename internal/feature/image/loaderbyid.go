package image

import (
	"context"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/google/uuid"
	"time"
)

type LoaderByIDKey struct {
	ID uuid.UUID
}

func NewConfiguredLoaderByID(repo *Repository, maxBatch int) *LoaderByID {
	return NewLoaderByID(LoaderByIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderByIDKey) ([]*Image, []error) {
			items := make([]*Image, len(keys))
			errors := make([]error, len(keys))

			ids := getUniqueImageIDs(keys)

			images, err := repo.Find(
				context.Background(),
				FindFilterIDAnyOf(ids),
				FindOrderByCreatedAt(false),
				condition.Limit(len(ids)),
			)
			if err != nil {
				for i := range keys {
					errors[i] = err
				}
			}

			groups := groupImagesByID(images)
			for i, key := range keys {
				if p, ok := groups[key.ID]; ok {
					items[i] = &p
				}
			}

			return items, errors
		},
	})
}

func getUniqueImageIDs(keys []LoaderByIDKey) []uuid.UUID {
	mapping := make(map[uuid.UUID]bool)

	for _, key := range keys {
		mapping[key.ID] = true
	}

	ids := make([]uuid.UUID, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupImagesByID(images []Image) map[uuid.UUID]Image {
	groups := make(map[uuid.UUID]Image)

	for _, p := range images {
		groups[p.ID] = p
	}

	return groups
}
