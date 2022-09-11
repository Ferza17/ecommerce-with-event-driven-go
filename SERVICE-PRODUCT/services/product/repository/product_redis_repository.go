package repository

import "github.com/go-redis/redis/v8"

type productRedisRepository struct {
	client *redis.Client
}

func NewProductRedisRepository(client *redis.Client) ProductRedisRepositoryStore {
	return &productRedisRepository{
		client: client,
	}
}
