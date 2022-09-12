package mongodb

import (
	"context"
	"time"

	"github.com/RoseRocket/xerrs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
	cartSchema "github.com/Ferza17/event-driven-cart-service/model/bson"
	"github.com/Ferza17/event-driven-cart-service/model/pb"
	"github.com/Ferza17/event-driven-cart-service/utils"
)

func (q *cartMongoDBRepository) CreateCart(ctx context.Context, session mongo.Session, request *pb.CreateCartRequest) (response *pb.Cart, err error) {
	var (
		now = time.Now().UTC()
	)
	response = &pb.Cart{}
	span, ctx := tracing.StartSpanFromContext(ctx, "CartMongoDBRepository-CreateCart")
	defer span.Finish()
	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		userCollection := q.db.Database(databaseCart).Collection(collectionCarts)
		result, err := userCollection.InsertOne(ctx, cartSchema.Cart{
			TotalPrice: 0,
			UserId:     request.GetUserId(),
			CreatedAt:  primitive.Timestamp{T: uint32(now.Unix()), I: 0},
			UpdatedAt:  primitive.Timestamp{T: uint32(now.Unix()), I: 0},
		})
		if err != nil {
			err = xerrs.Mask(err, utils.ErrInternalServerError)
			return err
		}
		response.Id = result.InsertedID.(primitive.ObjectID).Hex()
		response.UserId = request.GetUserId()
		return nil
	})
	return
}

func (q *cartMongoDBRepository) DeleteCartById(ctx context.Context, session mongo.Session, id string) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CartMongoDBRepository-DeleteCartById")
	defer span.Finish()
	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		primitiveId, _ := primitive.ObjectIDFromHex(id)
		userCollection := q.db.Database(databaseCart).Collection(collectionCarts)
		if err = userCollection.FindOneAndDelete(ctx, bson.M{"_id": primitiveId}).Err(); err != nil {
			return xerrs.Mask(err, utils.ErrInternalServerError)
		}
		return nil
	})
	return
}
