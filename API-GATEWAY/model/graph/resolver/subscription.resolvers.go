package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Ferza17/event-driven-api-gateway/model/graph/generated"
	"github.com/Ferza17/event-driven-api-gateway/model/graph/model"
	userPresenter "github.com/Ferza17/event-driven-api-gateway/module/user/presenter/graphql"
)

// SubscribeChangeCartState is the resolver for the subscribeChangeCartState field.
func (r *subscriptionResolver) SubscribeChangeCartState(ctx context.Context, input *model.SubscribeChangeCartState) (<-chan *model.Cart, error) {
	panic(fmt.Errorf("not implemented: SubscribeChangeCartState - subscribeChangeCartState"))
}

// SubscribeChangeUserState is the resolver for the subscribeChangeUserState field.
func (r *subscriptionResolver) SubscribeChangeUserState(ctx context.Context, input *model.SubscribeChangeUserState) (<-chan *model.User, error) {
	return userPresenter.HandleSubscribeChangeUserState(ctx, input)
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
