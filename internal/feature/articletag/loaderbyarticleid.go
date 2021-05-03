package articletag

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type LoaderByArticleIDKey struct {
	ArticleID uuid.UUID
}

func NewConfiguredLoaderByArticleID(repo *Repository, maxBatch int) *LoaderByArticleID {
	return NewLoaderByArticleID(LoaderByArticleIDConfig{
		Wait:     2 * time.Millisecond,
		MaxBatch: maxBatch,
		Fetch: func(keys []LoaderByArticleIDKey) ([][]ArticleTag, []error) {
			items := make([][]ArticleTag, len(keys))
			errors := make([]error, len(keys))

			ids := getUniqueArticleIDs(keys)

			articleTags, err := repo.Find(
				context.Background(),
				FindFilterArticleIDAnyOf(ids),
				FindOrderBySortRank(false),
			)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
			}

			groups := groupArticleTagsByArticleID(articleTags)
			for i, key := range keys {
				if at, ok := groups[key.ArticleID]; ok {
					items[i] = at
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

func groupArticleTagsByArticleID(articleTags []ArticleTag) map[uuid.UUID][]ArticleTag {
	groups := make(map[uuid.UUID][]ArticleTag)

	for _, at := range articleTags {
		groups[at.ArticleID] = append(groups[at.ArticleID], at)
	}

	return groups
}