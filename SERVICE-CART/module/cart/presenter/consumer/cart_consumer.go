package consumer

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/RoseRocket/xerrs"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
	"github.com/Ferza17/event-driven-cart-service/middleware"
	"github.com/Ferza17/event-driven-cart-service/model/pb"
	"github.com/Ferza17/event-driven-cart-service/module/cart"
	"github.com/Ferza17/event-driven-cart-service/saga"
	"github.com/Ferza17/event-driven-cart-service/utils"
)

type cartConsumerPresenter struct {
}

func NewCartConsumerPresenter() *cartConsumerPresenter {
	return &cartConsumerPresenter{}
}

func (c *cartConsumerPresenter) Consume(ctx context.Context, ch *amqp.Channel) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		c.consumeNewCartQueue(ctx, ch)
	}()
	go func() {
		defer wg.Done()
		c.consumeNewUserSAGAQueue(ctx, ch)
	}()
	wg.Wait()
}

func (c *cartConsumerPresenter) consumeNewCartQueue(ctx context.Context, ch *amqp.Channel) {
	var (
		cartUseCase = cart.GetCartUseCaseFromContext(ctx)
		tracer      = middleware.GetTracerFromContext(ctx)
		request     *pb.CreateCartRequest
	)
	opentracing.SetGlobalTracer(tracer)
	stopChan := make(chan bool)
	q, err := ch.QueueDeclare(
		string(utils.NewCartQueue),
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
			span, ctx := tracing.StartSpanFromContext(ctx, "consumeNewCartQueue")
			ctx = opentracing.ContextWithSpan(ctx, span)
			err = json.Unmarshal(d.Body, &request)
			log.Println("consumeNewCartQueue")
			_, err = cartUseCase.CreateCart(ctx, request)
			d.Ack(false)
			span.Finish()
		}
	}()
	<-stopChan
}

func (c *cartConsumerPresenter) consumeNewUserSAGAQueue(ctx context.Context, ch *amqp.Channel) {
	var (
		cartUseCase = cart.GetCartUseCaseFromContext(ctx)
		tracer      = middleware.GetTracerFromContext(ctx)
		request     saga.Step
	)
	opentracing.SetGlobalTracer(tracer)
	stopChan := make(chan bool)
	q, err := ch.QueueDeclare(
		string(utils.NewUserSagaQueue),
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
			span, ctx := tracing.StartSpanFromContext(ctx, "consumeNewUserSAGAQueue")
			ctx = opentracing.ContextWithSpan(ctx, span)
			err = json.Unmarshal(d.Body, &request)
			log.Println("consumeNewUserSAGAQueue")
			if request.Status == utils.SagaStatusFailed {
				err = cartUseCase.RollbackNewUserSAGA(ctx, request.TransactionId)
			}
			d.Ack(false)
			span.Finish()
		}
	}()
	<-stopChan
}
