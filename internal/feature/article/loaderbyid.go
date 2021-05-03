package article

import (
	"context"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/google/uuid"
	"time"
)

type LoaderByIDRepository interface {
	Find(ctx context.Context, conditions ...condition.Condition) ([]Article, error)
}

type LoaderByIDKey struct {
	ID uuid.UUID
}

func NewConfiguredLoaderByID(repo LoaderByIDRepository, maxBatch int) *LoaderByID {
	return NewLoaderByID(LoaderByIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderByIDKey) ([]*Article, []error) {
			items := make([]*Article, len(keys))
			errors := make([]error, len(keys))

			ctx, cancel := context.WithTimeout(
				context.Background(),
				50*time.Millisecond*time.Duration(len(keys)),
			)
			defer cancel()

			ids := getUniqueArticleIDs(keys)

			articles, err := repo.Find(
				ctx,
				FindFilterIDAnyOf(ids),
				FindOrderByCreatedAt(false),
				condition.Limit(len(ids)),
			)
			if err != nil {
				for i := range keys {
					errors[i] = err
				}
			}

			groups := groupArticlesByID(articles)
			for i, key := range keys {
				if a, ok := groups[key.ID]; ok {
					items[i] = &a
				}
			}

			return items, errors
		},
	})
}

func getUniqueArticleIDs(keys []LoaderByIDKey) []uuid.UUID {
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

func groupArticlesByID(articles []Article) map[uuid.UUID]Article {
	groups := make(map[uuid.UUID]Article)

	for _, a := range articles {
		groups[a.ID] = a
	}

	return groups
}
