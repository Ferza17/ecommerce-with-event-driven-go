package graphql

import (
	"context"

	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/model/graph/model"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/module/product"
)

func HandleFindProductById(ctx context.Context, id string) (response *model.Product, err error) {
	var (
		productUseCase = product.GetProductUseCaseFromContext(ctx)
	)
	response = &model.Product{}
	span, ctx := tracing.StartSpanFromContext(ctx, "ProductGRPCPresenter-HandleFindProductById")
	defer span.Finish()
	product, err := productUseCase.FindProductById(
		ctx,
		&pb.FindProductByIdRequest{
			Id: id,
		},
	)
	if err != nil {
		return
	}
	response = &model.Product{
		ID:          product.GetId(),
		Name:        product.GetName(),
		Description: product.GetDescription(),
		Uom:         product.GetUom(),
		Image:       product.GetImage(),
		Price:       product.GetPrice(),
		Stock:       int(product.GetStock()),
		CreatedAt:   int(product.GetCreatedAt()),
		UpdatedAt:   int(product.GetUpdatedAt()),
		DiscardedAt: int(product.GetDiscardedAt()),
	}
	return
}

func HandleFindProducts(ctx context.Context, input *model.FindProductsRequest) (response []*model.Product, err error) {
	var (
		productUseCase = product.GetProductUseCaseFromContext(ctx)
	)
	response = []*model.Product{}
	span, ctx := tracing.StartSpanFromContext(ctx, "ProductGRPCPresenter-HandleFindProducts")
	defer span.Finish()
	products, err := productUseCase.FindProducts(
		ctx,
		&pb.FindProductsRequest{
			ProductIds:  input.Ids,
			ProductName: input.Names,
			Limit:       int64(input.Limit),
			Page:        int64(input.Page),
		},
	)
	if err != nil {
		return
	}

	for _, product := range products.GetProducts() {
		response = append(response, &model.Product{
			ID:          product.GetId(),
			Name:        product.GetName(),
			Description: product.GetDescription(),
			Uom:         product.GetUom(),
			Image:       product.GetImage(),
			Price:       product.GetPrice(),
			Stock:       int(product.GetStock()),
			CreatedAt:   int(product.GetCreatedAt()),
			UpdatedAt:   int(product.GetUpdatedAt()),
			DiscardedAt: int(product.GetDiscardedAt()),
		})
	}
	return
}
