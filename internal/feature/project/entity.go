package project

import (
	"github.com/google/uuid"
	"time"
)

type Project struct {
	tableName struct{} `pg:"project"`

	CreatedAt  time.Time  `pg:",notnull,use_zero"`
	DeletedAt  *time.Time
	ID         uuid.UUID  `pg:"type:uuid,pk"`
	ModifiedAt time.Time  `pg:",notnull,use_zero"`
	Name       string     `pg:",notnull,use_zero"`
	Version    uint       `pg:",notnull,use_zero"`
}
