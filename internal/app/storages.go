package app

import (
	"github.com/acelot/articles/internal/feature/image"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/url"
	"strings"
)

type Storages struct {
	ImageStorage *image.Storage
}

func NewStorages(imageStorageURI string) (*Storages, error) {
	imageStorage, err := newImageStorage(imageStorageURI)
	if err != nil {
		return nil, err
	}

	return &Storages{
		ImageStorage: imageStorage,
	}, nil
}

func newImageStorage(uri string) (*image.Storage, error) {
	parsed, _ := url.Parse(uri)

	endpoint := parsed.Host
	bucketName := strings.Trim(parsed.Path, "/")
	id := parsed.User.Username()
	secret, _ := parsed.User.Password()
	useSSL := parsed.Scheme == "https"

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(id, secret, ""),
		Secure: useSSL,
	})

	if err != nil {
		return nil, err
	}

	return image.NewStorage(client, bucketName), nil
}