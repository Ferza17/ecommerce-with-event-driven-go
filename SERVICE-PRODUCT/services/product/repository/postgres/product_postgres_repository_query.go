package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/RoseRocket/xerrs"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Ferza17/event-driven-product-service/helper/tracing"
	"github.com/Ferza17/event-driven-product-service/model/pb"
	"github.com/Ferza17/event-driven-product-service/utils"
)

func (q *productPostgresRepository) FindProducts(ctx context.Context, request *pb.FindProductsRequest) (response *pb.FindProductsResponse, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "ProductSqlRepository-FindProductByID")
	defer span.Finish()
	return
}

func (q *productPostgresRepository) FindProductById(ctx context.Context, id string) (response *pb.Product, err error) {
	var (
		query = fmt.Sprintf(
			`SELECT 
    			id
			    , "name"
			    , description
			    , image
     			, stock
     			, price
     			, uom
			    , created_at
			    , updated_at
				, COALESCE(discarded_at, 0)	
    		FROM %s
			WHERE id = $1`, tableProduct)

		createdAt   int64
		updatedAt   int64
		discardedAt int64
	)
	response = &pb.Product{}
	if err = q.dbRead.QueryRowxContext(ctx, query, id).Scan(
		&response.Id,
		&response.Name,
		&response.Description,
		&response.Image,
		&response.Stock,
		&response.Price,
		&response.Uom,
		&createdAt,
		&updatedAt,
		&discardedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			err = xerrs.Mask(err, utils.ErrNotFound)
			return
		}
		err = xerrs.Mask(err, utils.ErrQueryRead)
	}
	response.CreatedAt = timestamppb.New(time.Unix(createdAt, 0))
	response.UpdatedAt = timestamppb.New(time.Unix(updatedAt, 0))
	response.DiscardedAt = timestamppb.New(time.Unix(discardedAt, 0))
	return
}
