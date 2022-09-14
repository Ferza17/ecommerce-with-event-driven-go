package middleware

import (
	"context"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Ferza17/event-driven-api-gateway/utils"
)

func RegisterRabbitMQAmqpHTTPContext(conn *amqp.Connection) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			if conn != nil {
				ctx = context.WithValue(ctx, string(utils.RabbitmqAmqpContextKey), conn)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func GetRabbitMQAmqpFromContext(ctx context.Context) *amqp.Connection {
	return ctx.Value(utils.RabbitmqAmqpContextKey).(*amqp.Connection)
}
