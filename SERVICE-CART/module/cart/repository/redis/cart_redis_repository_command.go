package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/RoseRocket/xerrs"

	"github.com/Ferza17/event-driven-cart-service/helper/tracing"
	"github.com/Ferza17/event-driven-cart-service/model/pb"
	"github.com/Ferza17/event-driven-cart-service/utils"
)

func (c *cartRedisRepository) SetCartLastStateByTransactionId(ctx context.Context, transactionId string, payload *pb.Cart) (err error) {
	var (
		now = time.Now().UTC()
	)
	span, ctx := tracing.StartSpanFromContext(ctx, "CartRedisRepository-SetCartLastStateByTransactionId")
	defer span.Finish()
	value, err := json.Marshal(payload)
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
		return
	}
	nextMinute := now.Add(time.Hour*time.Duration(0) +
		time.Minute*time.Duration(5) +
		time.Second*time.Duration(0))
	ttl := nextMinute.Sub(now).Seconds()
	err = c.client.Set(ctx, c.getRedisKey(cacheCartByTransactionId, transactionId), value, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
	}
	return
}

func (c *cartRedisRepository) DeleteCartLastStateByTransactionId(ctx context.Context, transactionId string) (err error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "CartRedisRepository-DeleteCartLastStateByTransactionId")
	defer span.Finish()
	if err = c.client.Del(ctx, c.getRedisKey(cacheCartByTransactionId, transactionId)).Err(); err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
	}
	return
}
