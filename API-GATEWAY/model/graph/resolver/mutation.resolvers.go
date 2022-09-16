package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Ferza17/event-driven-api-gateway/model/graph/generated"
	"github.com/Ferza17/event-driven-api-gateway/model/graph/model"
	userPresenter "github.com/Ferza17/event-driven-api-gateway/module/user/presenter/graphql"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input *model.RegisterRequest) (*model.CommandResponse, error) {
	return userPresenter.HandleUserRegister(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
