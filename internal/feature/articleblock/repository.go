package articleblock

import (
	"context"
	"errors"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

const articleIDConstraint string = "article_block_article_id_fkey"
const articleIDSortRankConstraint string = "article_block_article_id_sort_rank_key"

var NoRowsAffectedError = errors.New("no rows affected error")
var ArticleIDConstraintError = errors.New("article ID constraint error")
var ArticleIDSortRankConstraintError = errors.New("article ID and sort rank sort rank error")

type Repository struct {
	primaryDB   *pg.DB
	secondaryDB *pg.DB
}

type FindFilterArticleIDAnyOf []uuid.UUID
type FindFilterIDAnyOf []uuid.UUID
type FindFilterSortRankFrom string
type FindFilterSortRankTo string
type FindOrderBySortRank bool
type FindOrderByCreatedAt bool

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

func (r *Repository) Find(ctx context.Context, conditions ...condition.Condition) ([]ArticleBlock, error) {
	var items []ArticleBlock

	query := r.secondaryDB.ModelContext(ctx, &items)

	condition.Apply(query, conditions...)

	err := query.Select()

	return items, err
}

func (r *Repository) FindOneByID(ctx context.Context, id uuid.UUID) (*ArticleBlock, error) {
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

func (r *Repository) Count(ctx context.Context, estimateThreshold int, conditions ...condition.Condition) (int, error) {
	query := r.secondaryDB.ModelContext(ctx, (*ArticleBlock)(nil))

	condition.Apply(query, conditions...)

	if estimateThreshold > 0 {
		return query.CountEstimate(estimateThreshold)
	}

	return query.Count()
}

func (r *Repository) Insert(ctx context.Context, entity ArticleBlock) error {
	_, err := r.primaryDB.ModelContext(ctx, &entity).Insert()

	return specifyError(err)
}

func (r *Repository) Update(ctx context.Context, entity ArticleBlock) error {
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

	if constraint == articleIDSortRankConstraint {
		return ArticleIDSortRankConstraintError
	}

	return err
}