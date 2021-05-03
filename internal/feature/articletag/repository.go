package articletag

import (
	"context"
	"errors"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

const articleIDConstraint string = "article_tag_article_id_fkey"
const articleIDTagIDConstraint string = "article_tag_article_id_tag_id_key"
const articleIDSortRankConstraint string = "article_tag_article_id_sort_rank_key"
const tagIDConstraint string = "article_tag_tag_id_fkey"

var ArticleIDConstraintError = errors.New("article ID constraint error")
var ArticleIDSortRankConstraintError = errors.New("article ID and sort rank constraint error")
var ArticleIDTagIDConstraintError = errors.New("article ID and tag ID constraint error")
var NoRowsAffectedError = errors.New("no rows affected error")
var TagIDConstraintError = errors.New("tag ID constraint error")

type Repository struct {
	primaryDB   *pg.DB
	secondaryDB *pg.DB
}

type FindFilterArticleIDAnyOf []uuid.UUID
type FindFilterIDAnyOf []uuid.UUID
type FindFilterSortRankFrom string
type FindFilterSortRankTo string
type FindOrderBySortRank bool

func (uuids FindFilterArticleIDAnyOf) Apply(query *pg.Query) {
	query.Where(`article_id = ANY(?)`, pg.Array(uuids))
}

func (uuids FindFilterIDAnyOf) Apply(query *pg.Query) {
	query.Where(`id = ANY(?)`, pg.Array(uuids))
}

func (rank FindFilterSortRankFrom) Apply(query *pg.Query) {
	query.Where(`sort_rank >= ?`, rank)
}

func (rank FindFilterSortRankTo) Apply(query *pg.Query) {
	query.Where(`sort_rank <= ?`, rank)
}

func (isDesc FindOrderBySortRank) Apply(query *pg.Query) {
	dir := "ASC"
	if isDesc {
		dir = "DESC"
	}

	query.Order("sort_rank " + dir)
}

func NewRepository(primaryDB *pg.DB, secondaryDB *pg.DB) *Repository {
	return &Repository{
		primaryDB,
		secondaryDB,
	}
}

func (r *Repository) Find(ctx context.Context, conditions ...condition.Condition) ([]ArticleTag, error) {
	var items []ArticleTag

	query := r.secondaryDB.ModelContext(ctx, &items)

	condition.Apply(query, conditions...)

	err := query.Select()

	return items, err
}

func (r *Repository) FindOneByID(ctx context.Context, id uuid.UUID) (*ArticleTag, error) {
	entities, err := r.Find(
		ctx,
		FindFilterIDAnyOf{id},
		condition.Limit(1),
	)
	if err != nil {
		return nil, err
	}

	if len(entities) == 0 {
		return nil, nil
	}

	return &entities[0], nil
}

func (r *Repository) Count(ctx context.Context, conditions ...condition.Condition) (int, error) {
	query := r.secondaryDB.ModelContext(ctx, (*ArticleTag)(nil))

	condition.Apply(query, conditions...)

	return query.Count()
}

func (r *Repository) Insert(ctx context.Context, entity ArticleTag) error {
	_, err := r.primaryDB.ModelContext(ctx, &entity).Insert()

	return specifyError(err)
}

func (r *Repository) Update(ctx context.Context, entity ArticleTag) error {
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

	if constraint == articleIDConstraint {
		return ArticleIDConstraintError
	}

	if constraint == tagIDConstraint {
		return TagIDConstraintError
	}

	if constraint == articleIDTagIDConstraint {
		return ArticleIDTagIDConstraintError
	}

	if constraint == articleIDSortRankConstraint {
		return ArticleIDSortRankConstraintError
	}

	return err
}