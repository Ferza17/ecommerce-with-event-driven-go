package repository

import "github.com/go-redis/redis/v8"

type userRedisRepository struct {
	client *redis.Client
}

func NewUserRedisRepository(client *redis.Client) UserRedisRepositoryStore {
	return &userRedisRepository{
		client: client,
	}
}
