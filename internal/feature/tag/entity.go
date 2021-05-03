package tag

import (
	"github.com/google/uuid"
	"time"
)

type Tag struct {
	tableName struct{} `pg:"tag"`

	CreatedAt  time.Time  `pg:",notnull,use_zero"`
	DeletedAt  *time.Time
	ID         uuid.UUID  `pg:"type:uuid,pk"`
	ModifiedAt time.Time  `pg:",notnull,use_zero"`
	Name       string     `pg:",notnull,use_zero"`
	Version    uint       `pg:",notnull,use_zero"`
}
