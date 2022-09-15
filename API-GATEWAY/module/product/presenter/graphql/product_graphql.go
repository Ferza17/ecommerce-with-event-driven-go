package graphql

import (
	"github.com/graphql-go/graphql"

	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/module/product"
)

func HandleFindProductById(p graphql.ResolveParams) (response *pb.Product, err error) {
	var (
		ctx            = p.Context
		productUseCase = product.GetProductUseCaseFromContext(ctx)
	)
	span, ctx := tracing.StartSpanFromContext(ctx, "ProductGRPCPresenter-HandleFindProductById")
	defer span.Finish()
	response, err = productUseCase.FindProductById(
		ctx,
		&pb.FindProductByIdRequest{
			Id: p.Args["id"].(string),
		},
	)
	return
}

func HandleFindProducts(p graphql.ResolveParams) (response *pb.FindProductsResponse, err error) {
	var (
		ctx            = p.Context
		productUseCase = product.GetProductUseCaseFromContext(ctx)
	)
	span, ctx := tracing.StartSpanFromContext(ctx, "ProductGRPCPresenter-HandleFindProducts")
	defer span.Finish()
	response, err = productUseCase.FindProducts(
		ctx,
		&pb.FindProductsRequest{
			ProductIds:  p.Args["productIds"].([]string),
			ProductName: p.Args["productName"].([]string),
			Limit:       p.Args["limit"].(int64),
			Page:        p.Args["page"].(int64),
		},
	)
	return
}
