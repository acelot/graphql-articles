package dbmigration

import (
	"time"
)

type DBMigration struct {
	tableName struct{} `pg:"_db_migration"`

	AppliedAt *time.Time `pg:",use_zero"`
	Name      string `pg:",pk"`
}
