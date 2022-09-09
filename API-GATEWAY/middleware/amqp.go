package middleware

import (
	"context"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	rabbitmqAmqpContextKey = "rabbitmq_amqp_key"
)

func RegisterRabbitMQAmqpHTTPContext(conn *amqp.Connection) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			if conn != nil {
				ctx = context.WithValue(ctx, rabbitmqAmqpContextKey, conn)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func GetRabbitMQAmqpAccess(ctx context.Context) *amqp.Connection {
	return ctx.Value(rabbitmqAmqpContextKey).(*amqp.Connection)
}
