package publisher

import (
	"context"

	"github.com/RoseRocket/xerrs"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/utils"
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
