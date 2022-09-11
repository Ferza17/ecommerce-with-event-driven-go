package rabbitmq

import (
	"context"
	"sync"

	"github.com/RoseRocket/xerrs"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Ferza17/event-driven-cart-service/utils"
)

func Consumer(ctx context.Context, conn *amqp.Connection) {
	_, err := conn.Channel()
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}
	// Register Consumer
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
