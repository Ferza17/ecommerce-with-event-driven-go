package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type (
	redisKey            string
	cartRedisRepository struct {
		client *redis.Client
	}
)

const (
	cacheCartByTransactionId redisKey = "transaction-id-%s"
)

func NewCartRedisRepository(client *redis.Client) CartRedisRepositoryStore {
	return &cartRedisRepository{
		client: client,
	}
}

func (c *cartRedisRepository) getRedisKey(redisKey redisKey, uniqueStr string) string {
	return fmt.Sprintf(string(redisKey), uniqueStr)
}
