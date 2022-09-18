package mongodb

import (
	"context"
	"time"

	"github.com/RoseRocket/xerrs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	userSchema "github.com/Ferza17/event-driven-account-service/model/bson"
	"github.com/Ferza17/event-driven-account-service/model/pb"
	"github.com/Ferza17/event-driven-account-service/utils"
)

func (q *userMongoDBRepository) CreateUser(ctx context.Context, session mongo.Session, request *pb.RegisterRequest) (response *pb.User, err error) {
	var (
		now = time.Now().UTC()
	)
	response = &pb.User{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserMongoDBRepository-CreateUser")
	defer span.Finish()
	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		userCollection := q.db.Database(databaseUser).Collection(collectionUsers)
		result, err := userCollection.InsertOne(ctx, userSchema.User{
			Username:  request.GetUsername(),
			Email:     request.GetEmail(),
			Password:  request.GetPassword(),
			Devices:   nil,
			CreatedAt: primitive.Timestamp{T: uint32(now.Unix()), I: 0},
			UpdatedAt: primitive.Timestamp{T: uint32(now.Unix()), I: 0},
		})
		if err != nil {
			err = xerrs.Mask(err, utils.ErrInternalServerError)
			return err
		}
		response.Id = result.InsertedID.(primitive.ObjectID).Hex()
		response.Username = request.GetUsername()
		response.Email = request.GetEmail()
		response.Password = request.GetPassword()
		return nil
	})
	return
}

func (q *userMongoDBRepository) DeleteUserByUserId(ctx context.Context, session mongo.Session, id string) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserMongoDBRepository-DeleteUserByUserId")
	defer span.Finish()
	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		primitiveId, _ := primitive.ObjectIDFromHex(id)
		userCollection := q.db.Database(databaseUser).Collection(collectionUsers)
		if err = userCollection.FindOneAndDelete(ctx, bson.M{"_id": primitiveId}).Err(); err != nil {
			return xerrs.Mask(err, utils.ErrInternalServerError)
		}
		return nil
	})
	return
}

func (q *userMongoDBRepository) UpdateUserByUserId(ctx context.Context, session mongo.Session, request *pb.UpdateUserByUserIdRequest) (err error) {
	var (
		now = time.Now().UTC()
	)
	span, ctx := tracing.StartSpanFromContext(ctx, "UserMongoDBRepository-UpdateUserByUserId")
	defer span.Finish()
	err = mongo.WithSession(ctx, session, func(sessionContext mongo.SessionContext) error {
		primitiveId, err := primitive.ObjectIDFromHex(request.GetId())
		if err != nil {
			err = xerrs.Mask(err, utils.ErrBadRequest)
			return err
		}
		userCollection := q.db.Database(databaseUser).Collection(collectionUsers)

		update := bson.M{
			"$set": bson.M{
				"username":  request.GetUsername(),
				"email":     request.GetEmail(),
				"updatedAt": primitive.Timestamp{T: uint32(now.Unix()), I: 0},
			},
		}

		if _, err = userCollection.UpdateOne(ctx, bson.M{"_id": primitiveId}, update); err != nil {
			err = xerrs.Mask(err, utils.ErrInternalServerError)
			return err
		}
		return nil
	})
	return
}
