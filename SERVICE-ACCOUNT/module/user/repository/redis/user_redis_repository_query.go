package redis

import (
	"context"
	"encoding/json"

	"github.com/RoseRocket/xerrs"
	"github.com/go-redis/redis/v8"

	"github.com/Ferza17/event-driven-account-service/helper/tracing"
	"github.com/Ferza17/event-driven-account-service/model/pb"
	"github.com/Ferza17/event-driven-account-service/utils"
)

func (c *userRedisRepository) FindUserLastStateByTransactionId(ctx context.Context, transactionId string) (response *pb.User, err error) {
	response = &pb.User{}
	span, ctx := tracing.StartSpanFromContext(ctx, "UserRedisRepository-FindUserLastStateByTransactionId")
	defer span.Finish()
	result, err := c.client.Get(ctx, c.getRedisKey(cacheUserByTransactionId, transactionId)).Result()
	if err == redis.Nil {
		err = nil
		return
	}
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
	}
	if result == "" {
		return
	}
	if err = json.Unmarshal([]byte(result), &response); err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
	}
	return
}
