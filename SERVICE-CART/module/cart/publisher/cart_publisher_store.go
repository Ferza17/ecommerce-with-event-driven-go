package publisher

import (
	"context"

	"github.com/Ferza17/event-driven-cart-service/saga"
	"github.com/Ferza17/event-driven-cart-service/utils"
)

type CartPublisherStore interface {
	PublishOrdinaryMessage(ctx context.Context, queue utils.Event, payload string) (err error)
	PublishSagaMessage(ctx context.Context, sagaQueue utils.EventSaga, payload *saga.Step) (err error)
}
