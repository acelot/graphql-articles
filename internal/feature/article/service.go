package article

import (
	"context"
	"errors"
	"fmt"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/google/uuid"
	"time"
)

var NotFoundError = errors.New("article not found error")
var VersionMismatchError = errors.New("version mismatch error")

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository}
}

func (s *Service) FindArticles(
	filter *model.ArticleFindFilterInput,
	sort model.ArticleFindSortEnum,
	pageSize uint,
	pageNumber uint,
) ([]Article, error) {
	orderBy, err := mapFindSortEnumToRepositoryCondition(sort)
	if err != nil {
		return []Article{}, err
	}

	conditions := mapFindFilterInputToRepositoryConditions(filter)

	conditions = append(
		conditions,
		condition.Limit(pageSize),
		condition.Offset((pageNumber-1)*pageSize),
		orderBy,
	)

	return s.repository.Find(context.Background(), conditions...)
}

func (s *Service) CountArticles(filter *model.ArticleFindFilterInput, estimate uint) (uint, error) {
	conditions := mapFindFilterInputToRepositoryConditions(filter)

	c, err := s.repository.Count(context.Background(), int(estimate), conditions...)

	return uint(c), err
}

func (s *Service) CreateArticle(projectID uuid.UUID) (*Article, error) {
	createdAt := time.Now()

	entity := Article{
		CoverImageID: nil,
		CreatedAt: createdAt,
		ID: uuid.New(),
		ModifiedAt: createdAt,
		ProjectID: projectID,
		Title: fmt.Sprintf("New article %s", createdAt.Format("2006-01-02T15:04:05-0700")),
		Version: 0,
	}

	if err := s.repository.Insert(context.Background(), entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (s *Service) UpdateArticle(input model.ArticleUpdateInput) (*Article, error) {
	entity, err := s.repository.FindOneByID(context.Background(), input.ID)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, NotFoundError
	}

	entity.CoverImageID = input.CoverImageID
	entity.ModifiedAt = time.Now()
	entity.Title = input.Title
	entity.Version = input.Version

	err = s.repository.Update(context.Background(), entity)

	if errors.Is(err, NoRowsAffectedError) {
		return nil, VersionMismatchError
	}

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func mapFindFilterInputToRepositoryConditions(filter *model.ArticleFindFilterInput) (conditions []condition.Condition) {
	if filter != nil && filter.IDAnyOf != nil {
		conditions = append(conditions, FindFilterIDAnyOf(filter.IDAnyOf))
	}

	return
}

func mapFindSortEnumToRepositoryCondition(sort model.ArticleFindSortEnum) (condition.Condition, error) {
	switch sort {
	case model.ArticleFindSortEnumCreatedAtAsc:
		return FindOrderByCreatedAt(false), nil
	case model.ArticleFindSortEnumCreatedAtDesc:
		return FindOrderByCreatedAt(true), nil
	default:
		return nil, fmt.Errorf(`not mapped sort value "%s"`, sort)
	}
}