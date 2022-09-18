package subscriber

import (
	"context"

	"github.com/Ferza17/event-driven-api-gateway/model/graph/model"
)

type UserSubscriberStore interface {
	SubscribeNewUserState(ctx context.Context, id string) (userCh <-chan *model.User, err error)
}
