package middleware

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

const (
	rabbitmqAmqpContextKey = "rabbitmq_amqp_key"
)

func RegisterRabbitMQAmqpContext(conn *amqp.Connection, parentContext context.Context) (ctx context.Context) {
	ctx = context.WithValue(parentContext, rabbitmqAmqpContextKey, conn)
	return
}

func UnaryRegisterRabbitMQAmqpContext(conn *amqp.Connection) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		newCtx := context.WithValue(ctx, rabbitmqAmqpContextKey, conn)
		resp, err = handler(newCtx, req)
		return resp, err
	}
}

func GetRabbitMQAmqpAccess(ctx context.Context) *amqp.Connection {
	return ctx.Value(rabbitmqAmqpContextKey).(*amqp.Connection)
}
