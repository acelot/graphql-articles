package model

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
)

const pageSizeMin uint = 1
const pageSizeMax uint = 100

type PageSize uint

func MarshalPageSize(i uint) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatUint(uint64(i), 10))
	})
}

func UnmarshalPageSize(v interface{}) (uint, error) {
	parsed, err := UnmarshalUInt(v)
	if err != nil {
		return 0, err
	}

	if parsed < pageSizeMin {
		return 0, fmt.Errorf("value must be greater or equal than %d", pageSizeMin)
	}

	if parsed > pageSizeMax {
		return 0, fmt.Errorf("value must be less or equal than %d", pageSizeMax)
	}

	return parsed, nil
}
