package mongodb

import (
	"context"
	"time"

	"github.com/RoseRocket/xerrs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	userSchema "github.com/Ferza17/event-driven-account-service/model/bson"
	"github.com/Ferza17/event-driven-account-service/model/pb"
	"github.com/Ferza17/event-driven-account-service/utils"
)

func (q *userMongoDBRepository) FindUserByEmail(ctx context.Context, email string) (response *pb.User, err error) {
	var (
		rawUser userSchema.User
	)
	response = &pb.User{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserMongoDBRepository-FindUserByEmail")
	defer span.Finish()
	userCollection := q.db.Database(databaseUser).Collection(collectionUsers)
	result := userCollection.FindOne(ctx, bson.M{"email": email})
	if err = result.Decode(&rawUser); err != nil {
		if err == mongo.ErrNoDocuments {
			err = xerrs.Mask(err, utils.ErrNotFound)
			return
		}
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}
	response.Id = rawUser.Id.Hex()
	response.Username = rawUser.Username
	response.Email = rawUser.Email
	response.Password = rawUser.Password
	response.CreatedAt = timestamppb.New(time.Unix(int64(rawUser.CreatedAt.T), 0))
	response.UpdatedAt = timestamppb.New(time.Unix(int64(rawUser.UpdatedAt.T), 0))
	response.DiscardedAt = timestamppb.New(time.Unix(int64(rawUser.DiscardedAt.T), 0))
	for _, device := range rawUser.Devices {
		response.Devices = append(response.Devices, &pb.Device{
			Id:          device.Id.Hex(),
			DeviceId:    device.DeviceID,
			AccessToken: device.AccessToken,
			CreatedAt:   timestamppb.New(time.Unix(int64(device.CreatedAt.T), 0)),
			UpdatedAt:   timestamppb.New(time.Unix(int64(device.UpdatedAt.T), 0)),
			DiscardedAt: timestamppb.New(time.Unix(int64(device.DiscardedAt.T), 0)),
		})
	}
	return
}

func (q *userMongoDBRepository) FindUserById(ctx context.Context, id string) (response *pb.User, err error) {
	var (
		rawUser userSchema.User
	)
	response = &pb.User{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserMongoDBRepository-FindUserById")
	defer span.Finish()
	userCollection := q.db.Database(databaseUser).Collection(collectionUsers)
	primitiveUserId, err := primitive.ObjectIDFromHex(id)
	result := userCollection.FindOne(ctx, bson.M{"_id": primitiveUserId})
	if err = result.Decode(&rawUser); err != nil {
		if err == mongo.ErrNoDocuments {
			err = xerrs.Mask(err, utils.ErrNotFound)
			return
		}
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}
	response.Id = rawUser.Id.Hex()
	response.Username = rawUser.Username
	response.Email = rawUser.Email
	response.Password = rawUser.Password
	response.CreatedAt = timestamppb.New(time.Unix(int64(rawUser.CreatedAt.T), 0))
	response.UpdatedAt = timestamppb.New(time.Unix(int64(rawUser.UpdatedAt.T), 0))
	response.DiscardedAt = timestamppb.New(time.Unix(int64(rawUser.DiscardedAt.T), 0))
	for _, device := range rawUser.Devices {
		response.Devices = append(response.Devices, &pb.Device{
			Id:          device.Id.Hex(),
			DeviceId:    device.DeviceID,
			AccessToken: device.AccessToken,
			CreatedAt:   timestamppb.New(time.Unix(int64(device.CreatedAt.T), 0)),
			UpdatedAt:   timestamppb.New(time.Unix(int64(device.UpdatedAt.T), 0)),
			DiscardedAt: timestamppb.New(time.Unix(int64(device.DiscardedAt.T), 0)),
		})
	}
	return
}
