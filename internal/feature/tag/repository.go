package tag

import (
	"context"
	"errors"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

const tagNameConstraint string = "tag_name_key"

var TagNameConstraintError = errors.New("tag name constraint error")
var NoRowsAffectedError = errors.New("no rows affected error")

type Repository struct {
	primaryDB   *pg.DB
	secondaryDB *pg.DB
}

type FindFilterIDAnyOf []uuid.UUID
type FindOrderByName bool
type FindOrderByCreatedAt bool

func (uuids FindFilterIDAnyOf) Apply(query *pg.Query) {
	query.Where(`id = ANY(?)`, pg.Array(uuids))
}

func (isDesc FindOrderByName) Apply(query *pg.Query) {
	dir := "ASC"
	if isDesc {
		dir = "DESC"
	}

	query.Order("name " + dir)
}

func (isDesc FindOrderByCreatedAt) Apply(query *pg.Query) {
	dir := "ASC"
	if isDesc {
		dir = "DESC"
	}

	query.Order("created_at " + dir)
}

func NewRepository(primaryDB *pg.DB, secondaryDB *pg.DB) *Repository {
	return &Repository{
		primaryDB,
		secondaryDB,
	}
}

func (r *Repository) Find(ctx context.Context, conditions ...condition.Condition) ([]Tag, error) {
	var items []Tag

	query := r.secondaryDB.ModelContext(ctx, &items)

	condition.Apply(query, conditions...)

	err := query.Select()

	return items, err
}

func (r *Repository) FindOneByID(ctx context.Context, id uuid.UUID) (*Tag, error) {
	entities, err := r.Find(ctx, FindFilterIDAnyOf{id})
	if err != nil {
		return nil, err
	}

	if len(entities) == 0 {
		return nil, nil
	}

	return &entities[0], nil
}

func (r *Repository) Count(ctx context.Context, estimateThreshold int, conditions ...condition.Condition) (int, error) {
	query := r.secondaryDB.ModelContext(ctx, (*Tag)(nil))

	condition.Apply(query, conditions...)

	if estimateThreshold > 0 {
		return query.CountEstimate(estimateThreshold)
	}

	return query.Count()
}

func (r *Repository) Insert(ctx context.Context, entity Tag) error {
	_, err := r.primaryDB.ModelContext(ctx, &entity).Insert()

	return specifyError(err)
}

func (r *Repository) Update(ctx context.Context, entity Tag) error {
	currentVersion := entity.Version
	entity.Version++

	res, err := r.primaryDB.ModelContext(ctx, &entity).
		WherePK().
		Where("version = ?", currentVersion).
		Update()

	if err != nil {
		return specifyError(err)
	}

	if res.RowsAffected() == 0 {
		return NoRowsAffectedError
	}

	return nil
}

func specifyError(err error) error {
	pgErr, ok := err.(pg.Error)
	if !ok {
		return err
	}

	constraint := pgErr.Field([]byte("n")[0])

	if constraint == tagNameConstraint {
		return TagNameConstraintError
	}

	return err
}
