package publisher

import (
	"context"
	"encoding/json"

	"github.com/RoseRocket/xerrs"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	"github.com/Ferza17/event-driven-account-service/saga"
	"github.com/Ferza17/event-driven-account-service/utils"
)

type userPublisher struct {
	rabbitMQConnection *amqp.Connection
}

func NewUserPublisher(rabbitMQConnection *amqp.Connection) UserPublisherStore {
	return &userPublisher{
		rabbitMQConnection: rabbitMQConnection,
	}
}

func (p *userPublisher) PublishOrdinaryMessage(ctx context.Context, queue utils.Event, payload string) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserPublisher-PublishOrdinaryMessage")
	defer span.Finish()
	ch, err := p.rabbitMQConnection.Channel()
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}
	q, err := ch.QueueDeclare(
		string(queue),
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
	if _, err = ch.PublishWithDeferredConfirmWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		},
	); err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
	}
	return
}

func (p *userPublisher) PublishSagaMessage(ctx context.Context, sagaQueue utils.EventSaga, payload *saga.Step) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserPublisher-PublishSagaMessage")
	defer span.Finish()
	ch, err := p.rabbitMQConnection.Channel()
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}
	q, err := ch.QueueDeclare(
		string(sagaQueue),
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
	request, err := json.Marshal(payload)
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}
	if _, err = ch.PublishWithDeferredConfirmWithContext(
		ctx,
		"new-user-saga",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        request,
		},
	); err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
	}
	return
}

func (p *userPublisher) ParsePayloadToString(ctx context.Context, request interface{}) (response string, err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserPublisher-ParsePayloadToString")
	defer span.Finish()
	jsonString, err := json.Marshal(request)
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}
	response = string(jsonString)
	return
}
