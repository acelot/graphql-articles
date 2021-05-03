package articleblock

import (
	"context"
	"errors"
	"fmt"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/google/uuid"
	"github.com/xissy/lexorank"
	"time"
)

var ArticleNotFoundError = errors.New("article not found error")
var InvalidSortRankError = errors.New("invalid sort rank error")
var NotFoundError = errors.New("article block not found error")
var TypeMismatchError = errors.New("article block type mismatch error")
var VersionMismatchError = errors.New("version mismatch error")

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository}
}

func (s *Service) FindArticleBlocks(
	filter *model.ArticleBlockFindFilterInput,
	sort model.ArticleBlockFindSortEnum,
	pageSize uint,
	pageNumber uint,
) ([]ArticleBlock, error) {
	orderBy, err := mapFindSortEnumToRepositoryCondition(sort)
	if err != nil {
		return []ArticleBlock{}, err
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

func (s *Service) CountArticleBlocks(filter *model.ArticleBlockFindFilterInput, estimate uint) (uint, error) {
	conditions := mapFindFilterInputToRepositoryConditions(filter)

	c, err := s.repository.Count(context.Background(), int(estimate), conditions...)

	return uint(c), err
}

func (s *Service) CreateArticleBlock(input model.ArticleBlockCreateInput) (*ArticleBlock, error) {
	blockType, err := MapGqlArticleBlockTypeEnumToArticleBlockType(input.BlockType)
	if err != nil {
		return nil, err
	}

	newSortRank, _ := lexorank.Rank(input.SortRank.Prev, input.SortRank.Next)

	entity := ArticleBlock{
		ArticleID:  input.ArticleID,
		CreatedAt:  time.Now(),
		ID:         uuid.New(),
		ModifiedAt: time.Now(),
		SortRank:   newSortRank,
		Type:       blockType,
		Version:    0,
	}

	switch blockType {
	case ArticleBlockTypeHTML:
		entity.Data = marshalArticleBlockHTMLData(model.ArticleBlockHTMLDataInput{})
	case ArticleBlockTypeImage:
		entity.Data = marshalArticleBlockImageData(model.ArticleBlockImageDataInput{})
	default:
		return nil, fmt.Errorf(`unmapped ArticleBlockType "%s"`, blockType)
	}

	if err := s.repository.Insert(context.Background(), entity); err != nil {
		if errors.Is(err, ArticleIDConstraintError) {
			return nil, ArticleNotFoundError
		}

		if errors.Is(err, ArticleIDSortRankConstraintError) {
			return nil, InvalidSortRankError
		}

		return nil, err
	}

	return &entity, nil
}

func (s *Service) UpdateArticleBlock(input model.ArticleBlockUpdateInput) (*ArticleBlock, error) {
	entity, err := s.repository.FindOneByID(context.Background(), input.ID)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, NotFoundError
	}

	if input.Data.HTML != nil {
		if entity.Type != ArticleBlockTypeHTML {
			return nil, TypeMismatchError
		}

		entity.Data = marshalArticleBlockHTMLData(*input.Data.HTML)
	}

	if input.Data.Image != nil {
		if entity.Type != ArticleBlockTypeImage {
			return nil, TypeMismatchError
		}

		entity.Data = marshalArticleBlockImageData(*input.Data.Image)
	}

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

func (s *Service) MoveArticleBlock(input model.ArticleBlockMoveInput) (*ArticleBlock, error) {
	entity, err := s.repository.FindOneByID(context.Background(), input.ID)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, NotFoundError
	}

	newSortRank, _ := lexorank.Rank(input.SortRank.Prev, input.SortRank.Next)

	entity.SortRank = newSortRank
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

func mapFindFilterInputToRepositoryConditions(filter *model.ArticleBlockFindFilterInput) []condition.Condition {
	var conditions []condition.Condition

	if filter == nil {
		return conditions
	}

	if filter.IDAnyOf != nil && len(filter.IDAnyOf) > 0 {
		conditions = append(conditions, FindFilterIDAnyOf(filter.IDAnyOf))
	}

	return conditions
}

func mapFindSortEnumToRepositoryCondition(sort model.ArticleBlockFindSortEnum) (condition.Condition, error) {
	switch sort {
	case model.ArticleBlockFindSortEnumSortRankAsc:
		return FindOrderBySortRank(false), nil
	case model.ArticleBlockFindSortEnumSortRankDesc:
		return FindOrderBySortRank(true), nil
	case model.ArticleBlockFindSortEnumCreatedAtAsc:
		return FindOrderByCreatedAt(false), nil
	case model.ArticleBlockFindSortEnumCreatedAtDesc:
		return FindOrderByCreatedAt(true), nil
	default:
		return nil, fmt.Errorf("not mapped sort enum value %s", sort)
	}
}
