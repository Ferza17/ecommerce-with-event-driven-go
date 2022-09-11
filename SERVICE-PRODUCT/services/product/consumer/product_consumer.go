package consumer

import (
	"context"
	"log"

	"github.com/RoseRocket/xerrs"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Ferza17/event-driven-product-service/helper/tracing"
	"github.com/Ferza17/event-driven-product-service/middleware"
	"github.com/Ferza17/event-driven-product-service/utils"
)

func ConsumeQueueSAGA(ctx context.Context, ch *amqp.Channel) {
	var (
		//userUseCase = user.GetUserUseCaseFromContext(ctx)
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
