package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/RoseRocket/xerrs"

	"github.com/Ferza17/event-driven-product-service/helper/tracing"
	"github.com/Ferza17/event-driven-product-service/model/pb"
	"github.com/Ferza17/event-driven-product-service/utils"
)

func (q *productPostgresRepository) FindProducts(ctx context.Context, request *pb.FindProductsRequest) (response *pb.FindProductsResponse, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "ProductSqlRepository-FindProducts")
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
	)
	response = &pb.Product{}
	span, ctx := tracing.StartSpanFromContext(ctx, "ProductSqlRepository-FindProductById")
	defer span.Finish()
	if err = q.dbRead.QueryRowxContext(ctx, query, id).Scan(
		&response.Id,
		&response.Name,
		&response.Description,
		&response.Image,
		&response.Stock,
		&response.Price,
		&response.Uom,
		&response.CreatedAt,
		&response.UpdatedAt,
		&response.DiscardedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			err = xerrs.Mask(err, utils.ErrNotFound)
			return
		}
		err = xerrs.Mask(err, utils.ErrQueryRead)
	}
	return
}

func (q *productPostgresRepository) FindProductsByProductIds(ctx context.Context, ids []string) (response *pb.FindProductsByProductIdsResponse, err error) {
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
    		FROM %s`, tableProduct)
		placeholders []string
		params       []interface{}
	)
	response = &pb.FindProductsByProductIdsResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "ProductSqlRepository-FindProductsByProductIds")
	defer span.Finish()
	if len(ids) < 1 {
		return
	}
	for _, s := range ids {
		params = append(params, s)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(params)))
	}
	query += fmt.Sprintf(` WHERE id IN (%s) `, strings.Join(placeholders, ","))
	query += `ORDER BY id DESC`
	rows, err := q.dbRead.Query(query, params...)
	if err == sql.ErrNoRows {
		err = xerrs.Mask(err, utils.ErrNotFound)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var (
			product pb.Product
		)
		if err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Image,
			&product.Stock,
			&product.Price,
			&product.Uom,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DiscardedAt,
		); err != nil {
			err = xerrs.Mask(err, utils.ErrQueryRead)
			return
		}
		response.Products[product.Id] = &product
	}
	return
}
