package articleblock

import (
	"github.com/google/uuid"
	"time"
)

type ArticleBlockData map[string]interface{}

type ArticleBlock struct {
	tableName struct{} `pg:"article_block"`

	ArticleID  uuid.UUID        `pg:"type:uuid"`
	CreatedAt  time.Time        `pg:",notnull,use_zero"`
	Data       ArticleBlockData `pg:",notnull,use_zero"`
	DeletedAt  *time.Time
	ID         uuid.UUID        `pg:"type:uuid"`
	ModifiedAt time.Time        `pg:",notnull,use_zero"`
	SortRank   string           `pg:",notnull,use_zero"`
	Type       ArticleBlockType `pg:",notnull,use_zero"`
	Version    uint             `pg:",notnull,use_zero"`
}
