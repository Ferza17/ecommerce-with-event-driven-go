package consumer

import (
	"context"
	"log"
	"sync"

	"github.com/RoseRocket/xerrs"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
	"github.com/Ferza17/event-driven-cart-service/middleware"
	"github.com/Ferza17/event-driven-cart-service/utils"
)

func UserConsumerRegister(ctx context.Context, ch *amqp.Channel) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		ConsumeQueueNewUser(ctx, ch)
	}()
	go func() {
		defer wg.Done()
		ConsumeQueueSAGA(ctx, ch)
	}()
	wg.Wait()
}

func ConsumeQueueNewUser(ctx context.Context, ch *amqp.Channel) {
	var (
		//cartUseCase = cart.GetUserUseCaseFromContext(ctx)
		tracer = middleware.GetTracerFromContext(ctx)
	)
	opentracing.SetGlobalTracer(tracer)
	stopChan := make(chan bool)
	q, err := ch.QueueDeclare(
		UserQueueNewUser,
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
			span, ctx := tracing.StartSpanFromContext(ctx, "ConsumeQueueNewUser")
			ctx = opentracing.ContextWithSpan(ctx, span)
			//err = json.Unmarshal(d.Body, &request)
			log.Println("ConsumeQueueNewUser")
			d.Ack(true)
			span.Finish()
		}
	}()
	<-stopChan
}

func ConsumeQueueSAGA(ctx context.Context, ch *amqp.Channel) {
	var (
		//userUseCase = cart.GetUserUseCaseFromContext(ctx)
		tracer = middleware.GetTracerFromContext(ctx)
	)
	opentracing.SetGlobalTracer(tracer)
	stopChan := make(chan bool)
	q, err := ch.QueueDeclare(
		SAGAQueue,
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
			span, ctx := tracing.StartSpanFromContext(ctx, "ConsumeQueueSAGA")
			ctx = opentracing.ContextWithSpan(ctx, span)
			log.Println("ConsumeQueueSAGA")
			d.Ack(true)
			span.Finish()
		}
	}()
	<-stopChan
}
