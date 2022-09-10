package consumer

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/RoseRocket/xerrs"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	"github.com/Ferza17/event-driven-account-service/middleware"
	"github.com/Ferza17/event-driven-account-service/model/pb"
	"github.com/Ferza17/event-driven-account-service/services/user"
	"github.com/Ferza17/event-driven-account-service/utils"
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
		userUseCase = user.GetUserUseCaseFromContext(ctx)
		tracer      = middleware.GetTracerFromContext(ctx)
		request     *pb.RegisterRequest
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
			err = json.Unmarshal(d.Body, &request)
			log.Println("ConsumeQueueNewUser")
			_, err = userUseCase.CreateUser(ctx, request)
			d.Ack(true)
			span.Finish()
		}
	}()
	<-stopChan
}

func ConsumeQueueSAGA(ctx context.Context, ch *amqp.Channel) {
	var (
		//userUseCase = user.GetUserUseCaseFromContext(ctx)
		tracer  = middleware.GetTracerFromContext(ctx)
		request *pb.RegisterRequest
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
			err = json.Unmarshal(d.Body, &request)
			log.Println("ConsumeQueueSAGA")
			d.Ack(true)
			span.Finish()
		}
	}()
	<-stopChan
}
