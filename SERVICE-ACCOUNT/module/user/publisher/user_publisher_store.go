package publisher

import (
	"context"

	"github.com/Ferza17/event-driven-account-service/saga"
	"github.com/Ferza17/event-driven-account-service/utils"
)

type UserPublisherStore interface {
	PublishOrdinaryMessage(ctx context.Context, queue utils.Queue, payload string) (err error)
	PublishSagaMessage(ctx context.Context, sagaQueue utils.SagaQueue, payload *saga.Step) (err error)
}
