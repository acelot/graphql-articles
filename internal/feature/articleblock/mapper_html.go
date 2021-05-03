package articleblock

//goland:noinspection SpellCheckingInspection
import (
	"fmt"
	gqlmodel "github.com/acelot/articles/internal/gql/model"
	"github.com/mitchellh/mapstructure"
)

func mapBlockToGqlModelArticleBlockHTML(block ArticleBlock) (*gqlmodel.ArticleBlockHTML, error) {
	model := gqlmodel.ArticleBlockHTML{
		CreatedAt:  block.CreatedAt,
		DeletedAt:  block.DeletedAt,
		ID:         block.ID,
		ModifiedAt: block.ModifiedAt,
		SortRank:   block.SortRank,
		Version:    block.Version,
	}

	var err error

	model.Data, err = unmarshalArticleBlockHTMLData(block.Data)
	if err != nil {
		return nil, fmt.Errorf("cannot map block data, possible data corruption: %v", err)
	}

	return &model, nil
}

func unmarshalArticleBlockHTMLData(data map[string]interface{}) (*gqlmodel.ArticleBlockHTMLData, error) {
	parsed := &gqlmodel.ArticleBlockHTMLData{}

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

func marshalArticleBlockHTMLData(data gqlmodel.ArticleBlockHTMLDataInput) map[string]interface{} {
	result := map[string]interface {}{}

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: mapstructure.TextUnmarshallerHookFunc(),
		Result:     &result,
		TagName:    "json",
	})

	_ = decoder.Decode(data)

	return result
}