package publisher

import (
	"context"

	"github.com/Ferza17/event-driven-api-gateway/utils"
)

type UserPublisherStore interface {
	PublishOrdinaryMessage(ctx context.Context, event utils.Event, payload string) (err error)
}
