package image

import (
	"github.com/google/uuid"
	"time"
)

type Image struct {
	tableName struct{} `pg:"image"`

	CreatedAt   time.Time `pg:",notnull,use_zero"`
	DeletedAt   *time.Time
	Type        string    `pg:",notnull,use_zero"`
	Height      int       `pg:",notnull,use_zero"`
	ID          uuid.UUID `pg:"type:uuid,pk"`
	ModifiedAt  time.Time `pg:",notnull,use_zero"`
	Version     uint      `pg:",notnull,use_zero"`
	Width       int       `pg:",notnull,use_zero"`
}
