package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/model/schema"
)

type UserUseCaseStore interface {
	CreateUser(ctx context.Context, request *pb.RegisterRequest) (response schema.CommandResponse, err error)
	FindUserByEmailAndPassword(ctx context.Context, request *pb.LoginRequest) (response *pb.LoginResponse, err error)
	FindUserById(ctx context.Context, request *pb.FindUserByIdRequest) (response *pb.User, err error)
}
