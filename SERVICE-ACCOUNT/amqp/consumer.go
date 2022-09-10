package amqp

import (
	"context"
	"sync"

	"github.com/RoseRocket/xerrs"
	amqp "github.com/rabbitmq/amqp091-go"

	userConsumer "github.com/Ferza17/event-driven-account-service/services/user/consumer"
	"github.com/Ferza17/event-driven-account-service/utils"
)

func Consumer(ctx context.Context, conn *amqp.Connection) {
	ch, err := conn.Channel()
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}
	// Register Consumer
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		userConsumer.UserConsumerRegister(ctx, ch)
	}()
	wg.Wait()
}
