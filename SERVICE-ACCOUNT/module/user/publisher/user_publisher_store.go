package publisher

import (
	"context"

	"github.com/Ferza17/event-driven-account-service/saga"
	"github.com/Ferza17/event-driven-account-service/utils"
)

type UserPublisherStore interface {
	PublishOrdinaryMessage(ctx context.Context, queue utils.Event, payload string) (err error)
	PublishSagaMessage(ctx context.Context, sagaQueue utils.EventSaga, payload *saga.Step) (err error)
}
