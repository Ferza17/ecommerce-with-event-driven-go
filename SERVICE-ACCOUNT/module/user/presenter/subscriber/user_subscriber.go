package subscriber

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/RoseRocket/xerrs"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	"github.com/Ferza17/event-driven-account-service/middleware"
	"github.com/Ferza17/event-driven-account-service/model/pb"
	"github.com/Ferza17/event-driven-account-service/module/user"
	"github.com/Ferza17/event-driven-account-service/saga"
	"github.com/Ferza17/event-driven-account-service/utils"
)

type userSubscriberPresenter struct {
}

func NewUserSubscriberPresenter() *userSubscriberPresenter {
	return &userSubscriberPresenter{}
}

func (c *userSubscriberPresenter) Subscribe(ctx context.Context, ch *amqp.Channel) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		c.subscribeCreateUserEvent(ctx, ch)
	}()
	go func() {
		defer wg.Done()
		c.subscribeCreateCartEventSaga(ctx, ch)
	}()
	wg.Wait()
}

func (c *userSubscriberPresenter) subscribeCreateUserEvent(ctx context.Context, ch *amqp.Channel) {
	var (
		tracer      = middleware.GetTracerFromContext(ctx)
		userUseCase = user.GetUserUseCaseFromContext(ctx)
		request     pb.RegisterRequest
	)
	opentracing.SetGlobalTracer(tracer)
	stopChan := make(chan bool)
	q, err := ch.QueueDeclare(
		string(utils.CreateUserEvent),
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
			span, ctx := tracing.StartSpanFromContext(ctx, "subscribeCreateUserEvent")
			ctx = opentracing.ContextWithSpan(ctx, span)
			err = json.Unmarshal(d.Body, &request)
			if request.GetUsername() == "" {
				return
			}
			_, err = userUseCase.CreateUser(ctx, &request)
			d.Ack(false)
			span.Finish()
		}
	}()
	<-stopChan
}

func (c *userSubscriberPresenter) subscribeCreateCartEventSaga(ctx context.Context, ch *amqp.Channel) {
	var (
		userUseCase = user.GetUserUseCaseFromContext(ctx)
		tracer      = middleware.GetTracerFromContext(ctx)
		request     saga.Step
	)
	opentracing.SetGlobalTracer(tracer)
	stopChan := make(chan bool)
	q, err := ch.QueueDeclare(
		string(utils.CreateCartEventSaga),
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
			span, ctx := tracing.StartSpanFromContext(ctx, "subscribeCreateCartEventSaga")
			ctx = opentracing.ContextWithSpan(ctx, span)
			err = json.Unmarshal(d.Body, &request)
			if request.Status == utils.SagaStatusFailed {
				_ = userUseCase.RollbackNewUserSAGA(ctx, request.TransactionId)
			}
			d.Ack(false)
			span.Finish()
		}
	}()
	<-stopChan
}
