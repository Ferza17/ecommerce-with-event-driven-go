package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type (
	redisKey            string
	userRedisRepository struct {
		client *redis.Client
	}
)

const (
	cacheUserByTransactionId redisKey = "transaction-id-%s"
)

func NewUserRedisRepository(client *redis.Client) UserRedisRepositoryStore {
	return &userRedisRepository{
		client: client,
	}
}

func (c *userRedisRepository) getRedisKey(redisKey redisKey, uniqueStr string) string {
	return fmt.Sprintf(string(redisKey), uniqueStr)
}
