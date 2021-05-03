package model

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
)

type Html string

func MarshalHtml(h string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf(`"%s"`, h))
	})
}

func UnmarshalHtml(v interface{}) (string, error) {
	value, err := graphql.UnmarshalString(v)
	if err != nil {
		return "", err
	}

	// @TODO Validate HTML
	return value, nil
}
