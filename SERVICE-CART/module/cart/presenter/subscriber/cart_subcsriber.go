package subscriber

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/RoseRocket/xerrs"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
	"github.com/Ferza17/event-driven-cart-service/middleware"
	"github.com/Ferza17/event-driven-cart-service/model/pb"
	"github.com/Ferza17/event-driven-cart-service/module/cart"
	"github.com/Ferza17/event-driven-cart-service/utils"
)

type cartSubscriberPresenter struct {
}

func NewCartSubscriberPresenter() *cartSubscriberPresenter {
	return &cartSubscriberPresenter{}
}

func (c *cartSubscriberPresenter) Subscribe(ctx context.Context, ch *amqp.Channel) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.subscribeCreateCartEvent(ctx, ch)
	}()
	wg.Wait()
}

func (c *cartSubscriberPresenter) subscribeCreateCartEvent(ctx context.Context, ch *amqp.Channel) {
	var (
		cartUseCase = cart.GetCartUseCaseFromContext(ctx)
		tracer      = middleware.GetTracerFromContext(ctx)
		request     *pb.CreateCartRequest
	)
	opentracing.SetGlobalTracer(tracer)
	stopChan := make(chan bool)
	q, err := ch.QueueDeclare(
		string(utils.CrateCartEvent),
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}
	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}
	go func() {
		for d := range msgs {
			span, ctx := tracing.StartSpanFromContext(ctx, "subscribeCreateCartEvent")
			ctx = opentracing.ContextWithSpan(ctx, span)
			err = json.Unmarshal(d.Body, &request)
			_, err = cartUseCase.CreateCart(ctx, request)
			d.Ack(false)
			span.Finish()
		}
	}()
	<-stopChan
}
