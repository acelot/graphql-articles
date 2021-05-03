package project

import (
	"context"
	"errors"
	"fmt"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/google/uuid"
	"time"
)

var NotFoundError = errors.New("project not found error")
var VersionMismatchError = errors.New("version mismatch error")

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository}
}

func (s *Service) FindProjects(
	filter *model.ProjectFindFilterInput,
	sort model.ProjectFindSortEnum,
	pageSize uint,
	pageNumber uint,
) ([]Project, error) {
	orderBy, err := mapFindSortEnumToRepositoryCondition(sort)
	if err != nil {
		return []Project{}, err
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

func (s *Service) CountProjects(filter *model.ProjectFindFilterInput) (uint, error) {
	conditions := mapFindFilterInputToRepositoryConditions(filter)

	c, err := s.repository.Count(context.Background(), conditions...)

	return uint(c), err
}

func (s *Service) CreateProject(input model.ProjectCreateInput) (*Project, error) {
	entity := Project{
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

func (s *Service) UpdateProject(input model.ProjectUpdateInput) (*Project, error) {
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

func mapFindFilterInputToRepositoryConditions(filter *model.ProjectFindFilterInput) (conditions []condition.Condition) {
	if filter == nil {
		return
	}

	if filter.IDAnyOf != nil {
		conditions = append(conditions, FindFilterIDAnyOf(filter.IDAnyOf))
	}

	return
}

func mapFindSortEnumToRepositoryCondition(sort model.ProjectFindSortEnum) (condition.Condition, error) {
	switch sort {
	case model.ProjectFindSortEnumNameAsc:
		return FindOrderByName(false), nil
	case model.ProjectFindSortEnumNameDesc:
		return FindOrderByName(true), nil
	case model.ProjectFindSortEnumCreatedAtAsc:
		return FindOrderByCreatedAt(false), nil
	case model.ProjectFindSortEnumCreatedAtDesc:
		return FindOrderByCreatedAt(false), nil
	default:
		return condition.None{}, fmt.Errorf(`not mapped sort value "%s"`, sort)
	}
}
