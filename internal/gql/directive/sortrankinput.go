package directive

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/acelot/articles/internal/gql/model"
)

type SortRankInputDirectiveFunc = func(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver,
) (res interface{}, err error)

func NewSortRankInputDirective() SortRankInputDirectiveFunc {
	return func(
		ctx context.Context,
		obj interface{},
		next graphql.Resolver,
	) (res interface{}, err error) {
		inputObj, err := next(ctx)
		if err != nil {
			return inputObj, err
		}

		input, ok := inputObj.(*model.SortRankInput)
		if !ok {
			panic("@sortRankInput directive should only be used with SortRankInput input type")
		}

		if input.Prev < "0" || input.Prev >= "z" {
			return inputObj, errors.New("invalid prev value, must be in range of `[0-z)`")
		}

		if input.Next <= "0" || input.Next > "z" {
			return inputObj, errors.New("invalid next value, must be in range of `(0-z]`")
		}

		if input.Prev >= input.Next {
			return inputObj, errors.New("next value must be greater than prev value")
		}

		return inputObj, nil
	}
}
