package mongodb

import (
	"context"
	"time"

	"github.com/RoseRocket/xerrs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
	cartSchema "github.com/Ferza17/event-driven-cart-service/model/bson"
	"github.com/Ferza17/event-driven-cart-service/model/pb"
	"github.com/Ferza17/event-driven-cart-service/utils"
)

func (q *cartMongoDBRepository) FindCartByUserId(ctx context.Context, id string) (response *pb.Cart, err error) {
	var (
		rawCart cartSchema.Cart
	)
	response = &pb.Cart{}
	span, ctx := tracing.StartSpanFromContext(ctx, "CartMongoDBRepository-FindCartByUserId")
	defer span.Finish()
	userCollection := q.db.Database(databaseCart).Collection(collectionCarts)
	result := userCollection.FindOne(ctx, bson.M{"user_id": id})
	if err = result.Decode(&rawCart); err != nil {
		if err == mongo.ErrNoDocuments {
			err = xerrs.Mask(err, utils.ErrNotFound)
			return
		}
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}

	response.Id = rawCart.Id.Hex()
	response.UserId = rawCart.UserId
	response.TotalPrice = rawCart.TotalPrice
	response.CreatedAt = time.Unix(int64(rawCart.CreatedAt.T), 0).Unix()
	response.UpdatedAt = time.Unix(int64(rawCart.UpdatedAt.T), 0).Unix()
	response.DiscardedAt = time.Unix(int64(rawCart.DiscardedAt.T), 0).Unix()
	for _, item := range rawCart.CartItems {
		response.CartItems = append(response.CartItems, &pb.CartItem{
			Id:          item.Id.Hex(),
			ProductId:   item.ProductId,
			Quantity:    item.Quantity,
			Price:       item.Price,
			Note:        item.Note,
			CreatedAt:   time.Unix(int64(item.CreatedAt.T), 0).Unix(),
			UpdatedAt:   time.Unix(int64(item.UpdatedAt.T), 0).Unix(),
			DiscardedAt: time.Unix(int64(item.DiscardedAt.T), 0).Unix(),
		})
	}
	return
}
