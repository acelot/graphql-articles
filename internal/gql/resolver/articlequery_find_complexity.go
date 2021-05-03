package resolver

import "github.com/acelot/articles/internal/gql/model"

func ArticleQueryFindComplexity(childComplexity int, filter *model.ArticleFindFilterInput, sort model.ArticleFindSortEnum, pageSize uint, pageNumber uint) int {
	return 10 + int(pageSize)*childComplexity
}

func ArticleFindListTotalCountComplexity(childComplexity int, estimate uint) int {
	return 10 + childComplexity
}
