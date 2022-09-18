package subscriber

import (
	"context"
	"encoding/json"

	"github.com/RoseRocket/xerrs"
	"github.com/opentracing/opentracing-go"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/Ferza17/event-driven-api-gateway/helper/tracing"
	"github.com/Ferza17/event-driven-api-gateway/model/graph/model"
	"github.com/Ferza17/event-driven-api-gateway/model/pb"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

type userSubscriber struct {
	rabbitMQConnection *amqp.Connection
}

func NewUserSubscriber(
	rabbitMQConnection *amqp.Connection,
) UserSubscriberStore {
	return &userSubscriber{
		rabbitMQConnection: rabbitMQConnection,
	}
}

func (s *userSubscriber) SubscribeNewUserState(ctx context.Context, id string) (<-chan *model.User, error) {
	var (
		user *pb.User
	)
	userCh := make(chan *model.User)
	errCh := make(chan error)

	ch, err := s.rabbitMQConnection.Channel()
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return nil, err
	}
	q, err := ch.QueueDeclare(
		string(utils.UserNewState),
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return nil, err
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
		return nil, err
	}
	go func() {
		for d := range msgs {
			span, ctx := tracing.StartSpanFromContext(ctx, "SubscribeNewUserState")
			ctx = opentracing.ContextWithSpan(ctx, span)
			if err = json.Unmarshal(d.Body, &user); err != nil {
				err = xerrs.Mask(err, utils.ErrInternalServerError)
				errCh <- err
			}
			d.Ack(true)
			span.Finish()
			if user.Id == id && user != nil {
				userCh <- &model.User{
					ID:          user.Id,
					Username:    user.Username,
					Email:       user.Email,
					Password:    user.Password,
					CreatedAt:   int(user.CreatedAt),
					UpdatedAt:   int(user.UpdatedAt),
					DiscardedAt: int(user.DiscardedAt),
				}
			}
		}
	}()
	<-userCh
	return userCh, nil
}
