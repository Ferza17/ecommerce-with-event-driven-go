package rabbitmq

import (
	"context"
	"sync"

	"github.com/RoseRocket/xerrs"
	amqp "github.com/rabbitmq/amqp091-go"

	userConsumer "github.com/Ferza17/event-driven-account-service/module/user/presenter/subscriber"
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
		userConsumer.
			NewUserSubscriberPresenter().
			Subscribe(ctx, ch)
	}()
	wg.Wait()
}
