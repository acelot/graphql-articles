package model

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"io"
)

type Uuid uuid.UUID

func MarshalUuid(u uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf(`"%s"`, u.String()))
	})
}

func UnmarshalUuid(v interface{}) (uuid.UUID, error) {
	value, err := graphql.UnmarshalString(v)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.Parse(value)
}
