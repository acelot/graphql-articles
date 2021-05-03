package image

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
	"time"
)

type Storage struct {
	client     *minio.Client
	bucketName string
}

func NewStorage(client *minio.Client, bucketName string) *Storage {
	return &Storage{client, bucketName}
}

func (s *Storage) Store(ctx context.Context, id uuid.UUID, mimeType string, fileSize int64, reader io.Reader) error {
	_, err := s.client.PutObject(
		ctx,
		s.bucketName,
		id.String(),
		reader,
		fileSize,
		minio.PutObjectOptions{
			ContentType: mimeType,
		},
	)

	return err
}

func (s *Storage) MakeTempURL(ctx context.Context, id uuid.UUID, ttl time.Duration) (*url.URL, error) {
	return s.client.PresignedGetObject(ctx, s.bucketName, id.String(), ttl, url.Values{})
}