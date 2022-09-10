package repository

import (
	"context"
	"time"

	"github.com/RoseRocket/xerrs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"

	userSchema "github.com/Ferza17/event-driven-account-service/model/bson"

	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	"github.com/Ferza17/event-driven-account-service/model/pb"
	"github.com/Ferza17/event-driven-account-service/utils"
)

const (
	databaseUser    = "user"
	collectionUsers = "users"
)

type userNOSQLRepository struct {
	db *mongo.Client
}

func NewUserMongoDBRepository(db *mongo.Client) UserMongoDBRepositoryStore {
	return &userNOSQLRepository{
		db: db,
	}
}

func (q *userNOSQLRepository) CreateUser(ctx context.Context, session mongo.Session, request *pb.RegisterRequest) (response *pb.RegisterResponse, err error) {
	var (
		now = time.Now().UTC()
	)
	response = &pb.RegisterResponse{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserNOSQLRepository-CreateUser")
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
		response.UserId = result.InsertedID.(primitive.ObjectID).Hex()
		return nil
	})
	return
}

func (q *userNOSQLRepository) FindUserByEmail(ctx context.Context, email string) (response *pb.User, err error) {
	var (
		rawUser userSchema.User
	)
	response = &pb.User{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserNOSQLRepository-FindUserByEmail")
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

func (q *userNOSQLRepository) FindUserById(ctx context.Context, id string) (response *pb.User, err error) {
	var (
		rawUser userSchema.User
	)
	response = &pb.User{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserNOSQLRepository-FindUserByEmail")
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

func (q *userNOSQLRepository) CreateNoSQLSession() (session mongo.Session, err error) {
	session, err = q.db.StartSession()
	if err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxBegin)
	}
	return
}

func (q *userNOSQLRepository) StartTransaction(session mongo.Session) (err error) {
	if err = session.StartTransaction(); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxBegin)
	}
	return
}

func (q *userNOSQLRepository) AbortTransaction(session mongo.Session, ctx context.Context) (err error) {
	if err = session.AbortTransaction(ctx); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxRollback)
	}
	return
}

func (q *userNOSQLRepository) CommitTransaction(session mongo.Session, ctx context.Context) (err error) {
	if err = session.CommitTransaction(ctx); err != nil {
		err = xerrs.Mask(err, utils.ErrQueryTxCommit)
	}
	return
}
