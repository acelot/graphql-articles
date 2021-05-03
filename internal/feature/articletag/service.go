package articletag

import (
	"context"
	"errors"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/google/uuid"
	"github.com/xissy/lexorank"
	"time"
)

var ArticleNotFoundError = errors.New("article not found error")
var DuplicateTagError = errors.New("tag not found error")
var InvalidSortRankError = errors.New("invalid sort rank error")
var NotFoundError = errors.New("article tag not found error")
var TagNotFoundError = errors.New("tag not found error")
var VersionMismatchError = errors.New("version mismatch error")

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository}
}

func (s *Service) CreateArticleTag(input model.ArticleTagCreateInput) (*ArticleTag, error) {
	newSortRank, _ := lexorank.Rank(input.SortRank.Prev, input.SortRank.Next)

	entity := ArticleTag{
		ArticleID:  input.ArticleID,
		CreatedAt:  time.Now(),
		ID:         uuid.New(),
		ModifiedAt: time.Now(),
		SortRank:   newSortRank,
		TagID:      input.TagID,
		Version:    0,
	}

	err := s.repository.Insert(context.Background(), entity)

	if errors.Is(err, ArticleIDConstraintError) {
		return nil, ArticleNotFoundError
	}

	if errors.Is(err, ArticleIDTagIDConstraintError) {
		return nil, DuplicateTagError
	}

	if errors.Is(err, ArticleIDSortRankConstraintError) {
		return nil, InvalidSortRankError
	}

	if errors.Is(err, TagIDConstraintError) {
		return nil, TagNotFoundError
	}

	return &entity, nil
}

func (s *Service) MoveArticleTag(input model.ArticleTagMoveInput) (*ArticleTag, error) {
	entity, err := s.repository.FindOneByID(context.Background(), input.ID)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, NotFoundError
	}

	newSortRank, _ := lexorank.Rank(input.SortRank.Prev, input.SortRank.Next)

	entity.ModifiedAt = time.Now()
	entity.SortRank = newSortRank
	entity.Version = input.Version

	err = s.repository.Update(context.Background(), *entity)

	if errors.Is(err, ArticleIDSortRankConstraintError) {
		return nil, InvalidSortRankError
	}

	if errors.Is(err, NoRowsAffectedError) {
		return nil, VersionMismatchError
	}

	if err != nil {
		return nil, err
	}

	return entity, nil
}
