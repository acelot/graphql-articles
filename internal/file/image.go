package file

import (
	"errors"
	"github.com/h2non/bimg"
)

type ImageMeta struct {
	Type   string
	Width  int
	Height int
}

func GetImageMeta(buffer []byte) (*ImageMeta, error) {
	image := bimg.NewImage(buffer)

	imageType := image.Type()
	if imageType == "unknown" {
		return nil, errors.New("cannot recognize image type")
	}

	imageSize, err := image.Size()
	if err != nil {
		return nil, err
	}

	return &ImageMeta{
		Type:   imageType,
		Width:  imageSize.Width,
		Height: imageSize.Height,
	}, nil
}
