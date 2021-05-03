package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/acelot/articles/internal/feature/articleblock"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (r *subscriptionResolver) ArticleBlockCreated(ctx context.Context, articleID uuid.UUID) (<-chan model.ArticleBlockInterface, error) {
	isSocketClosed := false
	resultChan := make(chan model.ArticleBlockInterface, 1)

	go func() {
		<-ctx.Done()

		r.env.Logger.Debug("websocket is closed")

		isSocketClosed = true
	}()

	go func() {
	Loop:
		for {
			if isSocketClosed {
				break Loop
			}

			lastBlocks, err := r.env.Repositories.ArticleBlockRepository.Find(
				context.Background(),
				articleblock.FindFilterArticleIDAnyOf{articleID},
				articleblock.FindOrderByCreatedAt(true),
				condition.Limit(1),
			)

			if err != nil {
				r.env.Logger.Error("articleblock.Repository.Find", zap.Error(err))

				continue
			}

			if len(lastBlocks) == 0 {
				continue
			}

			gqlArticleBlock, err := articleblock.MapOneToGqlModel(lastBlocks[0])
			if err != nil {
				r.env.Logger.Error("articleblock.MapOneToGqlModel", zap.Error(err))

				continue
			}

			resultChan <- gqlArticleBlock

			time.Sleep(5 * time.Second)
		}
	}()

	return resultChan, nil
}
