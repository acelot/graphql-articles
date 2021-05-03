package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"github.com/acelot/articles/internal/gql/runtime"
)

// Subscription returns runtime.SubscriptionResolver implementation.
func (r *Resolver) Subscription() runtime.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
