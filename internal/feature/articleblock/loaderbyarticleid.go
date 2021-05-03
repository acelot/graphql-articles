package articleblock

import (
	"context"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/google/uuid"
	"time"
)

type LoaderByArticleIDRepository interface {
	Find(ctx context.Context, conditions ...condition.Condition) ([]ArticleBlock, error)
}

type LoaderByArticleIDKey struct {
	ArticleID uuid.UUID
}

func NewConfiguredLoaderByArticleID(repo LoaderByArticleIDRepository, maxBatch int) *LoaderByArticleID {
	return NewLoaderByArticleID(LoaderByArticleIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderByArticleIDKey) ([][]ArticleBlock, []error) {
			items := make([][]ArticleBlock, len(keys))
			errors := make([]error, len(keys))

			ids := getUniqueArticleIDs(keys)

			blocks, err := repo.Find(
				context.Background(),
				FindFilterArticleIDAnyOf(ids),
				FindOrderBySortRank(false),
			)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
			}

			groups := groupBlocksByArticleID(blocks)
			for i, key := range keys {
				if b, ok := groups[key.ArticleID]; ok {
					items[i] = b
				}
			}

			return items, errors
		},
	})
}

func getUniqueArticleIDs(keys []LoaderByArticleIDKey) []uuid.UUID {
	mapping := make(map[uuid.UUID]bool)

	for _, key := range keys {
		mapping[key.ArticleID] = true
	}

	ids := make([]uuid.UUID, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}

func groupBlocksByArticleID(blocks []ArticleBlock) map[uuid.UUID][]ArticleBlock {
	groups := make(map[uuid.UUID][]ArticleBlock)

	for _, b := range blocks {
		groups[b.ArticleID] = append(groups[b.ArticleID], b)
	}

	return groups
}