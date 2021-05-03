package resolver

import (
	"github.com/acelot/articles/internal/app"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	env *app.Env
}

func NewResolver(env *app.Env) *Resolver {
	return &Resolver{
		env,
	}
}

