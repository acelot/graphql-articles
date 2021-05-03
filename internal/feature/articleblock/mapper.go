package articleblock

import (
	"fmt"
	"github.com/acelot/articles/internal/gql/model"
)

type ArticleBlockType string

const (
	ArticleBlockTypeHTML  ArticleBlockType = "html"
	ArticleBlockTypeImage ArticleBlockType = "image"
)

func MapOneToGqlModel(block ArticleBlock) (model.ArticleBlockInterface, error) {
	switch block.Type {
	case ArticleBlockTypeHTML:
		return mapBlockToGqlModelArticleBlockHTML(block)
	case ArticleBlockTypeImage:
		return mapBlockToGqlModelArticleBlockImage(block)
	default:
		return nil, fmt.Errorf(`cannot map block with unknown type "%s"`, block.Type)
	}
}

func MapManyToGqlModels(blocks []ArticleBlock) ([]model.ArticleBlockInterface, error) {
	items := make([]model.ArticleBlockInterface, len(blocks))

	for i, entity := range blocks {
		model, err := MapOneToGqlModel(entity)
		if err != nil {
			return items, err
		}

		items[i] = model
	}

	return items, nil
}

func MapGqlArticleBlockTypeEnumToArticleBlockType(blockType model.ArticleBlockTypeEnum) (ArticleBlockType, error) {
	switch blockType {
	case model.ArticleBlockTypeEnumHTML:
		return ArticleBlockTypeHTML, nil
	case model.ArticleBlockTypeEnumImage:
		return ArticleBlockTypeImage, nil
	default:
		return "", fmt.Errorf(`unmapped ArticleBlockTypeEnum value "%s"`, blockType)
	}
}