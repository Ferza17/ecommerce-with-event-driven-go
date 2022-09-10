package middleware

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"

	"github.com/Ferza17/event-driven-account-service/utils"
)

func RegisterRabbitMQAmqpContext(conn *amqp.Connection, parentContext context.Context) (ctx context.Context) {
	ctx = context.WithValue(parentContext, utils.RabbitmqAmqpContextKey, conn)
	return
}

func UnaryRegisterRabbitMQAmqpContext(conn *amqp.Connection) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		newCtx := context.WithValue(ctx, utils.RabbitmqAmqpContextKey, conn)
		resp, err = handler(newCtx, req)
		return resp, err
	}
}

func GetRabbitMQAmqpFromContext(ctx context.Context) *amqp.Connection {
	return ctx.Value(utils.RabbitmqAmqpContextKey).(*amqp.Connection)
}
