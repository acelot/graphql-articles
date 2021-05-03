package articletag

import (
	"github.com/google/uuid"
	"time"
)

type ArticleTag struct {
	tableName struct{} `pg:"article_tag"`

	ArticleID  uuid.UUID `pg:"type:uuid,notnull"`
	CreatedAt  time.Time `pg:",notnull,use_zero"`
	ID         uuid.UUID `pg:"type:uuid,pk"`
	ModifiedAt time.Time `pg:",notnull,use_zero"`
	SortRank   string    `pg:",notnull,use_zero"`
	TagID      uuid.UUID `pg:"type:uuid,notnull"`
	Version    uint      `pg:",notnull,use_zero"`
}
