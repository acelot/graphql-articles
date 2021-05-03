package model

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
)

type UInt uint

func MarshalUInt(i uint) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatUint(uint64(i), 10))
	})
}

func UnmarshalUInt(v interface{}) (uint, error) {
	switch v := v.(type) {
	case string:
		u64, err := strconv.ParseUint(v, 10, 32)
		return uint(u64), err
	case int:
		return uint(v), nil
	case int64:
		return uint(v), nil
	case json.Number:
		u64, err := strconv.ParseUint(string(v), 10, 32)
		return uint(u64), err
	default:
		return 0, fmt.Errorf("%T is not an uint", v)
	}
}
