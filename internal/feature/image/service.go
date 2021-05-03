package image

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/acelot/articles/internal/file"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/acelot/articles/internal/gql/model"
	"github.com/google/uuid"
	"io"
	"net/url"
	"time"
)

var NotSupportedTypeError = errors.New("image type not supported error")
var NotRecognizedError = errors.New("image not recognized error")

type Service struct {
	repository *Repository
	storage    *Storage
}

func NewService(repository *Repository, storage *Storage) *Service {
	return &Service{repository, storage}
}

func (s *Service) FindImages(
	filter *model.ImageFindFilterInput,
	sort model.ImageFindSortEnum,
	pageSize uint,
	pageNumber uint,
) ([]Image, error) {
	orderBy, err := mapFindSortEnumToRepositoryCondition(sort)
	if err != nil {
		return []Image{}, err
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

func (s *Service) CountImages(filter *model.ImageFindFilterInput, estimate uint) (uint, error) {
	conditions := mapFindFilterInputToRepositoryConditions(filter)

	c, err := s.repository.Count(context.Background(), int(estimate), conditions...)

	return uint(c), err
}

func (s *Service) UploadImage(uploadedFile *graphql.Upload) (*Image, error) {
	imageBytes := make([]byte, uploadedFile.Size)

	bytesRead, err := io.ReadFull(uploadedFile.File, imageBytes)
	if err != nil {
		return nil, err
	}

	if int64(bytesRead) != uploadedFile.Size {
		return nil, errors.New("the number of bytes bytesRead doesn't match the content-length")
	}

	imageMeta, err := file.GetImageMeta(imageBytes)
	if err != nil {
		return nil, NotRecognizedError
	}

	if isImageTypeSupported(imageMeta.Type) == false {
		return nil, NotSupportedTypeError
	}

	imageID := uuid.New()
	createdAt := time.Now()

	if err := s.storage.Store(
		context.Background(),
		imageID,
		"image/" + imageMeta.Type,
		int64(bytesRead),
		bytes.NewReader(imageBytes),
	); err != nil {
		return nil, err
	}

	entity := Image{
		CreatedAt:  createdAt,
		DeletedAt:  nil,
		Height:     imageMeta.Height,
		ID:         imageID,
		ModifiedAt: createdAt,
		Type:       imageMeta.Type,
		Version:    0,
		Width:      imageMeta.Width,
	}

	if err := s.repository.Insert(context.Background(), entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (s *Service) GetImageDownloadURL(imageID uuid.UUID) (*url.URL, error) {
	return s.storage.MakeTempURL(context.Background(), imageID, time.Hour)
}

func isImageTypeSupported(imageType string) bool {
	if imageType == "jpeg" {
		return true
	}

	if imageType == "png" {
		return true
	}

	if imageType == "webp" {
		return true
	}

	if imageType == "avif" {
		return true
	}

	return false
}

func mapFindFilterInputToRepositoryConditions(filter *model.ImageFindFilterInput) (conditions []condition.Condition) {
	if filter == nil {
		return
	}

	if filter.IDAnyOf != nil {
		conditions = append(conditions, FindFilterIDAnyOf(filter.IDAnyOf))
	}

	return
}

func mapFindSortEnumToRepositoryCondition(sort model.ImageFindSortEnum) (condition.Condition, error) {
	switch sort {
	case model.ImageFindSortEnumCreatedAtAsc:
		return FindOrderByCreatedAt(false), nil
	case model.ImageFindSortEnumCreatedAtDesc:
		return FindOrderByCreatedAt(true), nil
	default:
		return condition.None{}, fmt.Errorf(`not mapped sort value "%s"`, sort)
	}
}
