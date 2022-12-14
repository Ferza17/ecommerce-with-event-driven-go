package middleware

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"

	"github.com/Ferza17/event-driven-account-service/utils"
)

func RegisterRabbitMQAmqpContext(conn *amqp.Connection, ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, utils.RabbitmqAmqpContextKey, conn)
	return ctx
}

func UnaryRegisterRabbitMQAmqpContext(conn *amqp.Connection) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(context.WithValue(ctx, utils.RabbitmqAmqpContextKey, conn), req)
	}
}

func GetRabbitMQAmqpFromContext(ctx context.Context) *amqp.Connection {
	return ctx.Value(utils.RabbitmqAmqpContextKey).(*amqp.Connection)
}
