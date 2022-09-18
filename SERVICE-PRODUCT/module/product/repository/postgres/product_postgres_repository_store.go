package postgres

import (
	"context"

	"github.com/Ferza17/event-driven-product-service/model/pb"
)

type ProductPostgresRepositoryCommand interface {
}

type ProductPostgresRepositoryQuery interface {
	FindProductsByProductIds(ctx context.Context, ids []string) (response *pb.FindProductsByProductIdsResponse, err error)
	FindProducts(ctx context.Context, request *pb.FindProductsRequest) (response *pb.FindProductsResponse, err error)
	FindProductById(ctx context.Context, id string) (response *pb.Product, err error)
}

type ProductPostgresRepositoryStore interface {
	ProductPostgresRepositoryCommand
	ProductPostgresRepositoryQuery
}
