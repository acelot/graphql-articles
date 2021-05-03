package dbmigration

import (
	"context"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Repository struct {
	db *pg.DB
}

type FindFilterNameAnyOf []string
type FindFilterIsAppliedOnly bool
type FindOrderByName bool
type FindOrderByAppliedAt bool

func (names FindFilterNameAnyOf) Apply(query *pg.Query) {
	query.Where(`name = ANY(?)`, pg.Array(names))
}

func (isAppliedOnly FindFilterIsAppliedOnly) Apply(query *pg.Query) {
	if isAppliedOnly {
		query.Where(`applied_at IS NOT NULL`)
	}
}

func (isDesc FindOrderByName) Apply(query *pg.Query) {
	dir := "ASC"
	if isDesc {
		dir = "DESC"
	}

	query.Order("name " + dir)
}

func (isDesc FindOrderByAppliedAt) Apply(query *pg.Query) {
	dir := "ASC"
	if isDesc {
		dir = "DESC"
	}

	query.Order("applied_at " + dir)
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		db,
	}
}

func (r *Repository) EnsureTable() error {
	return r.db.Model((*DBMigration)(nil)).CreateTable(
		&orm.CreateTableOptions{
			IfNotExists: true,
		},
	)
}

func (r *Repository) Find(ctx context.Context, conditions ...condition.Condition) ([]DBMigration, error) {
	var items []DBMigration

	query := r.db.ModelContext(ctx, &items)

	condition.Apply(query, conditions...)

	err := query.Select()

	return items, err
}

func (r *Repository) Create(m *DBMigration) error {
	_, err := r.db.Model(m).Insert()

	return err
}

func (r *Repository) Update(m *DBMigration) error {
	_, err := r.db.Model(m).WherePK().Update()

	return err
}
