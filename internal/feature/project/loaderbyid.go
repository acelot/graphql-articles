package project

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
		Fetch: func(keys []LoaderByIDKey) ([]*Project, []error) {
			items := make([]*Project, len(keys))
			errors := make([]error, len(keys))

			ids := getUniqueProjectIDs(keys)

			projects, err := repo.Find(
				context.Background(),
				FindFilterIDAnyOf(ids),
				condition.Limit(len(ids)),
				FindOrderByCreatedAt(false),
			)
			if err != nil {
				for i := range keys {
					errors[i] = err
				}
			}

			groups := groupProjectsByID(projects)
			for i, key := range keys {
				if p, ok := groups[key.ID]; ok {
					items[i] = &p
				}
			}

			return items, errors
		},
	})
}

func getUniqueProjectIDs(keys []LoaderByIDKey) []uuid.UUID {
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

func groupProjectsByID(projects []Project) map[uuid.UUID]Project {
	groups := make(map[uuid.UUID]Project)

	for _, p := range projects {
		groups[p.ID] = p
	}

	return groups
}
