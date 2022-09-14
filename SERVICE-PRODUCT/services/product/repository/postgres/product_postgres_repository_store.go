package postgres

import (
	"context"

	"github.com/Ferza17/event-driven-product-service/model/pb"
)

type ProductPostgresRepositoryCommand interface {
	FindProducts(ctx context.Context, request *pb.FindProductsRequest) (response *pb.FindProductsResponse, err error)
	FindProductById(ctx context.Context, id string) (response *pb.Product, err error)
}

type ProductPostgresRepositoryQuery interface {
}

type ProductPostgresRepositoryStore interface {
	ProductPostgresRepositoryCommand
	ProductPostgresRepositoryQuery
}
