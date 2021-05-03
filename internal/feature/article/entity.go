package article

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	tableName struct{} `pg:"article"`

	CoverImageID *uuid.UUID `pg:"type:uuid"`
	CreatedAt    time.Time  `pg:",use_zero"`
	DeletedAt    *time.Time
	ID           uuid.UUID  `pg:"type:uuid,pk"`
	ModifiedAt   time.Time  `pg:",use_zero"`
	ProjectID    uuid.UUID  `pg:"type:uuid"`
	Title        string     `pg:",use_zero"`
	Version      uint       `pg:",use_zero"`
}
