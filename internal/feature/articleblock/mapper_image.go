package articleblock

//goland:noinspection ALL
import (
	"fmt"
	gqlmodel "github.com/acelot/articles/internal/gql/model"
	"github.com/mitchellh/mapstructure"
)

func mapBlockToGqlModelArticleBlockImage(block ArticleBlock) (*gqlmodel.ArticleBlockImage, error) {
	model := gqlmodel.ArticleBlockImage{
		CreatedAt:  block.CreatedAt,
		DeletedAt:  block.DeletedAt,
		ID:         block.ID,
		ModifiedAt: block.ModifiedAt,
		SortRank:   block.SortRank,
		Version:    block.Version,
	}

	var err error

	model.Data, err = unmarshalArticleBlockImageData(block.Data)
	if err != nil {
		return nil, fmt.Errorf("cannot map block data, possible data corruption: %v", err)
	}

	return &model, nil
}

func unmarshalArticleBlockImageData(data map[string]interface{}) (*gqlmodel.ArticleBlockImageData, error) {
	parsed := &gqlmodel.ArticleBlockImageData{}

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: mapstructure.TextUnmarshallerHookFunc(),
		Result:     parsed,
		TagName:    "json",
	})

	err := decoder.Decode(data)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func marshalArticleBlockImageData(data gqlmodel.ArticleBlockImageDataInput) map[string]interface{} {
	result := map[string]interface {}{}

	_ = mapstructure.Decode(data, result)

	return result
}
