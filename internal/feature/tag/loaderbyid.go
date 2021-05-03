package tag

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
		Fetch: func(keys []LoaderByIDKey) ([]*Tag, []error) {
			items := make([]*Tag, len(keys))
			errors := make([]error, len(keys))

			ids := getUniqueTagIDs(keys)

			tags, err := repo.Find(context.Background(), FindFilterIDAnyOf(ids), condition.Limit(len(ids)))
			if err != nil {
				for i := range keys {
					errors[i] = err
				}
			}

			groups := groupTagsByID(tags)
			for i, key := range keys {
				if t, ok := groups[key.ID]; ok {
					items[i] = &t
				}
			}

			return items, errors
		},
	})
}

func getUniqueTagIDs(keys []LoaderByIDKey) []uuid.UUID {
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

func groupTagsByID(tags []Tag) map[uuid.UUID]Tag {
	groups := make(map[uuid.UUID]Tag)

	for _, t := range tags {
		groups[t.ID] = t
	}

	return groups
}
