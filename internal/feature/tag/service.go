package tag

import (
	"context"
	"errors"
	"fmt"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/google/uuid"
	"time"
)

var NotFoundError = errors.New("tag not found error")
var VersionMismatchError = errors.New("version mismatch error")

type Service struct {
	repository *Repository
}

func NewService(tagRepository *Repository) *Service {
	return &Service{tagRepository}
}

func (s *Service) FindTags(
	filter *model.TagFindFilterInput,
	sort model.TagFindSortEnum,
	pageSize uint,
	pageNumber uint,
) ([]Tag, error) {
	orderBy, err := mapFindSortEnumToRepositoryCondition(sort)
	if err != nil {
		return []Tag{}, err
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

func (s *Service) CountTags(filter *model.TagFindFilterInput, estimate uint) (uint, error) {
	conditions := mapFindFilterInputToRepositoryConditions(filter)

	c, err := s.repository.Count(context.Background(), int(estimate), conditions...)

	return uint(c), err
}

func (s *Service) CreateTag(input model.TagCreateInput) (*Tag, error) {
	entity := Tag{
		CreatedAt:  time.Now(),
		ID:         uuid.New(),
		ModifiedAt: time.Now(),
		Name:       input.Name,
		Version:    0,
	}

	if err := s.repository.Insert(context.Background(), entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (s *Service) UpdateTag(input model.TagUpdateInput) (*Tag, error) {
	entity, err := s.repository.FindOneByID(context.Background(), input.ID)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, NotFoundError
	}

	entity.ModifiedAt = time.Now()
	entity.Name = input.Name
	entity.Version = input.Version

	err = s.repository.Update(context.Background(), *entity)

	if errors.Is(err, NoRowsAffectedError) {
		return nil, VersionMismatchError
	}

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func mapFindFilterInputToRepositoryConditions(filter *model.TagFindFilterInput) (conditions []condition.Condition) {
	if filter == nil {
		return
	}

	if filter.IDAnyOf != nil {
		conditions = append(conditions, FindFilterIDAnyOf(filter.IDAnyOf))
	}

	return
}

func mapFindSortEnumToRepositoryCondition(sort model.TagFindSortEnum) (condition.Condition, error) {
	switch sort {
	case model.TagFindSortEnumNameAsc:
		return FindOrderByName(false), nil
	case model.TagFindSortEnumNameDesc:
		return FindOrderByName(true), nil
	case model.TagFindSortEnumCreatedAtAsc:
		return FindOrderByCreatedAt(false), nil
	case model.TagFindSortEnumCreatedAtDesc:
		return FindOrderByCreatedAt(true), nil
	default:
		return condition.None{}, fmt.Errorf(`not mapped sort value "%s"`, sort)
	}
}
