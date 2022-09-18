package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Ferza17/event-driven-api-gateway/model/graph/generated"
	"github.com/Ferza17/event-driven-api-gateway/model/graph/model"
	cartPresenter "github.com/Ferza17/event-driven-api-gateway/module/cart/presenter/graphql"
	productPresenter "github.com/Ferza17/event-driven-api-gateway/module/product/presenter/graphql"
	userPresenter "github.com/Ferza17/event-driven-api-gateway/module/user/presenter/graphql"
)

// FindUserByID is the resolver for the findUserById field.
func (r *queryResolver) FindUserByID(ctx context.Context) (*model.User, error) {
	return userPresenter.HandleFindUserByIdFindUserByID(ctx)
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, input *model.LoginRequest) (*model.LoginResponse, error) {
	return userPresenter.HandleUserLogin(ctx, input)
}

// FindCartByUserID is the resolver for the findCartByUserId field.
func (r *queryResolver) FindCartByUserID(ctx context.Context) (*model.Cart, error) {
	return cartPresenter.HandleFindCartByUserID(ctx)
}

// FindCartItemsArgsType is the resolver for the findCartItemsArgsType field.
func (r *queryResolver) FindCartItemsArgsType(ctx context.Context, input *model.FindCartItemsRequest) ([]*model.CartItem, error) {
	panic(fmt.Errorf("not implemented: FindCartItemsArgsType - findCartItemsArgsType"))
}

// FindProductByID is the resolver for the findProductById field.
func (r *queryResolver) FindProductByID(ctx context.Context, input string) (*model.Product, error) {
	return productPresenter.HandleFindProductById(ctx, input)
}

// FindProducts is the resolver for the findProducts field.
func (r *queryResolver) FindProducts(ctx context.Context, input *model.FindProductsRequest) ([]*model.Product, error) {
	return productPresenter.HandleFindProducts(ctx, input)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
