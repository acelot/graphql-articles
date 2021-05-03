package model

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"net/url"
)

type Url url.URL

func MarshalUrl(u url.URL) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf(`"%s"`, u.String()))
	})
}

func UnmarshalUrl(v interface{}) (url.URL, error) {
	value, err := graphql.UnmarshalString(v)
	if err != nil {
		return url.URL{}, err
	}

	parsedURL, err := url.ParseRequestURI(value)
	if err != nil {
		return url.URL{}, err
	}

	return *parsedURL, nil
}
