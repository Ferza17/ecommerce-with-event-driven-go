package usecase

import (
	"context"

	"github.com/Ferza17/event-driven-account-service/model/pb"
)

type UserUseCaseCommand interface {
	CreateUser(ctx context.Context, request *pb.RegisterRequest) (response *pb.RegisterResponse, err error)
	UpdateUserByUserId(ctx context.Context, request *pb.UpdateUserByUserIdRequest) (err error)
}

type UserUseCaseQuery interface {
	FindUserByEmailAndPassword(ctx context.Context, request *pb.LoginRequest) (response *pb.LoginResponse, err error)
	FindUserByEmail(ctx context.Context, request *pb.FindUserByEmailRequest) (response *pb.User, err error)
	FindUserById(ctx context.Context, request *pb.FindUserByIdRequest) (response *pb.User, err error)
}

type UserUseCaseCompensate interface {
	RollbackUserCrateNewUserSAGA(ctx context.Context, transactionId string) (err error)
}

type UserUseCaseStore interface {
	UserUseCaseCommand
	UserUseCaseQuery
	UserUseCaseCompensate
}
