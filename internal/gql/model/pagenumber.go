package model

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
)

const pageNumberMin uint = 1
const pageNumberMax uint = 100

type PageNumber uint

func MarshalPageNumber(i uint) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatUint(uint64(i), 10))
	})
}

func UnmarshalPageNumber(v interface{}) (uint, error) {
	parsed, err := UnmarshalUInt(v)
	if err != nil {
		return 0, err
	}

	if parsed < pageSizeMin {
		return 0, fmt.Errorf("value must be greater or equal than %d", pageNumberMin)
	}

	if parsed > pageSizeMax {
		return 0, fmt.Errorf("value must be less or equal than %d", pageNumberMax)
	}

	return parsed, nil
}
