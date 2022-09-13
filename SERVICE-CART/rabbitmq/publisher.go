package rabbitmq

import (
	"context"
	"sync"

	"github.com/RoseRocket/xerrs"
	amqp "github.com/rabbitmq/amqp091-go"

	cartConsumer "github.com/Ferza17/event-driven-cart-service/module/cart/presenter/subscriber"
	"github.com/Ferza17/event-driven-cart-service/utils"
)

func Subscriber(ctx context.Context, conn *amqp.Connection) {
	ch, err := conn.Channel()
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}
	// Register Subscriber
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		cartConsumer.NewCartSubscriberPresenter().Subscribe(ctx, ch)
	}()
	wg.Wait()
}
